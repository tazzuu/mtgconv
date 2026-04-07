package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"os"
)

// TopCommander is the small subset of fields we care about for a simple ranked list.
type TopCommander struct {
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	NumDecks int    `json:"num_decks"`
	URL      string `json:"url"`
}

// cardList is the structure EDHREC embeds for the initial commanders page.
type cardList struct {
	Header    string         `json:"header"`
	Tag       string         `json:"tag"`
	More      string         `json:"more"`
	Cardviews []TopCommander `json:"cardviews"`
}

// nextData is the JSON stored inside the page's __NEXT_DATA__ script tag.
//
// How we found this:
// 1. Fetch https://edhrec.com/commanders
// 2. Inspect the HTML source / DevTools and look for the <script id="__NEXT_DATA__"> blob
// 3. That blob contains the initial page payload, including the first 100 commanders
//
// EDHREC also exposes equivalent Next.js JSON at a URL like:
//   https://edhrec.com/_next/data/<buildId>/commanders.json
// but <buildId> changes over time, so parsing __NEXT_DATA__ from the HTML page is the
// more stable way to bootstrap the request flow.
type nextData struct {
	Props struct {
		PageProps struct {
			Data struct {
				Container struct {
					JSONDict struct {
						Cardlists []cardList `json:"cardlists"`
					} `json:"json_dict"`
				} `json:"container"`
			} `json:"data"`
		} `json:"pageProps"`
	} `json:"props"`
}

// paginatedCommandersPage is the shape returned by json.edhrec.com/pages/... for "load more".
type paginatedCommandersPage struct {
	Cardviews   []TopCommander `json:"cardviews"`
	IsPaginated bool           `json:"is_paginated"`
	More        string         `json:"more"`
}

// FetchTopCommanders retrieves the full ranked commander list currently shown on
// https://edhrec.com/commanders.
//
// What it does:
// 1. Download the HTML page
// 2. Extract the embedded __NEXT_DATA__ JSON bootstrap payload
// 3. Read the initial 100 commanders from cardlists[*].cardviews
// 4. Follow the "more" links at https://json.edhrec.com/pages/... until exhausted
// 5. Return a single flat slice of commanders in rank order
func FetchTopCommanders(ctx context.Context, client *http.Client) ([]TopCommander, error) {
	const (
		commandersPageURL = "https://edhrec.com/commanders"
		jsonPagesBaseURL  = "https://json.edhrec.com/pages/"
	)

	// Step 1: fetch the public commanders page HTML.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, commandersPageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("build commanders page request: %w", err)
	}
	req.Header.Set("User-Agent", "mtgconv-thank-you/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch commanders page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch commanders page: unexpected status %s", resp.Status)
	}

	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read commanders page HTML: %w", err)
	}
	html := string(htmlBytes)

	// Step 2: extract the Next.js bootstrap payload from:
	//   <script id="__NEXT_DATA__" type="application/json">...</script>
	//
	// This is the easiest stable entrypoint because it gives us the initial data
	// without needing to guess the current Next.js build ID.
	const nextDataPrefix = `<script id="__NEXT_DATA__" type="application/json">`
	start := strings.Index(html, nextDataPrefix)
	if start == -1 {
		return nil, fmt.Errorf("__NEXT_DATA__ script tag not found")
	}
	start += len(nextDataPrefix)

	endRel := strings.Index(html[start:], "</script>")
	if endRel == -1 {
		return nil, fmt.Errorf("__NEXT_DATA__ closing script tag not found")
	}

	nextDataJSON := html[start : start+endRel]

	var bootstrap nextData
	if err := json.Unmarshal([]byte(nextDataJSON), &bootstrap); err != nil {
		return nil, fmt.Errorf("decode __NEXT_DATA__: %w", err)
	}

	// Step 3: find the card list that represents the "Top Commanders" view.
	//
	// As verified on 2026-04-03, the page currently exposes a single list tagged
	// "past2years", and it contains ranks 1-100 plus a "more" field for pagination.
	cardlists := bootstrap.Props.PageProps.Data.Container.JSONDict.Cardlists
	if len(cardlists) == 0 {
		return nil, fmt.Errorf("no commander cardlists found in bootstrap payload")
	}

	var current cardList
	found := false
	for _, cl := range cardlists {
		if cl.Tag == "past2years" || strings.EqualFold(cl.Header, "Past 2 Years") {
			current = cl
			found = true
			break
		}
	}
	if !found {
		// Fallback: if EDHREC changes the tag/header, use the first list.
		current = cardlists[0]
	}

	// Start with the first page of commanders already embedded in the HTML.
	commanders := make([]TopCommander, 0, len(current.Cardviews)+200)
	commanders = append(commanders, current.Cardviews...)

	// Step 4: follow EDHREC's pagination chain.
	//
	// The initial payload includes a relative path like:
	//   commanders/year-past2years-1.json
	//
	// By inspecting the client bundle and network behavior, we can see EDHREC loads
	// that from:
	//   https://json.edhrec.com/pages/<more>
	//
	// Each subsequent page returns:
	//   {
	//     "cardviews": [...],
	//     "is_paginated": true|false,
	//     "more": "commanders/year-past2years-2.json"
	//   }
	nextMore := current.More
	for nextMore != "" {
		pageURL := jsonPagesBaseURL + nextMore

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
		if err != nil {
			return nil, fmt.Errorf("build paginated request for %q: %w", pageURL, err)
		}
		req.Header.Set("User-Agent", "mtgconv-example/1.0")

		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("fetch paginated commanders page %q: %w", pageURL, err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("fetch paginated commanders page %q: unexpected status %s", pageURL, resp.Status)
		}

		var page paginatedCommandersPage
		err = json.NewDecoder(resp.Body).Decode(&page)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("decode paginated commanders page %q: %w", pageURL, err)
		}

		commanders = append(commanders, page.Cardviews...)

		if !page.IsPaginated || page.More == "" {
			break
		}
		nextMore = page.More
	}

	return commanders, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	commanders, err := FetchTopCommanders(ctx, client)
	if err != nil {
		panic(err)
	}

	// Final output: a simple ranked list of commanders.
	// for _, c := range commanders {
	// 	fmt.Printf("%3d. %-40s %6d decks\n", c.Rank, c.Name, c.NumDecks)
	// }
	// Build the full ranked table as plain text.
	// We keep it in memory first so we can both print it and write it to a file.
	var table strings.Builder
	for _, c := range commanders {
		fmt.Fprintf(&table, "%3d. %-40s %6d decks\n", c.Rank, c.Name, c.NumDecks)
	}

	// Save the full table to disk.
	// Example output file:
	//   top_commanders_table.txt
	if err := os.WriteFile("top_commanders_table.txt", []byte(table.String()), 0644); err != nil {
		panic(fmt.Errorf("write full table file: %w", err))
	}

	// Also print the table to stdout, same as before.
	fmt.Print(table.String())

	// Build a second file containing only the top 200 commander names.
	// If fewer than 200 are returned, just write however many we have.
	limit := 200
	if len(commanders) < limit {
		limit = len(commanders)
	}

	var names strings.Builder
	for i := 0; i < limit; i++ {
		names.WriteString(commanders[i].Name)
		names.WriteByte('\n')
	}

	// Save the top-200 names to disk.
	// Example output file:
	//   top_200_commander_names.txt
	if err := os.WriteFile("top_200_commander_names.txt", []byte(names.String()), 0644); err != nil {
		panic(fmt.Errorf("write top 200 names file: %w", err))
	}

	fmt.Printf("wrote %d commanders to %s\n", len(commanders), "top_commanders_table.txt")
	fmt.Printf("wrote top %d names to %s\n", limit, "top_200_commander_names.txt")


}
