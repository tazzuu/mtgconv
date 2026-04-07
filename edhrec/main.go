package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// TopCommander is the small subset of fields we care about for a simple ranked list.
type TopCommander struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rank        int    `json:"rank"`
	NumDecks    int    `json:"num_decks"`
	URL         string `json:"url"`
	ScryfallURI string `json:"scryfall_uri"`
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
//
//	https://edhrec.com/_next/data/<buildId>/commanders.json
//
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
	BuildID string `json:"buildId"`
}

// paginatedCommandersPage is the shape returned by json.edhrec.com/pages/... for "load more".
type paginatedCommandersPage struct {
	Cardviews   []TopCommander `json:"cardviews"`
	IsPaginated bool           `json:"is_paginated"`
	More        string         `json:"more"`
}

// SnapshotManifest records the raw and normalized artifacts saved for one fetch run.
type SnapshotManifest struct {
	FetchedAt          string   `json:"fetched_at"`
	SnapshotDir        string   `json:"snapshot_dir"`
	CommandersURL      string   `json:"commanders_url"`
	HTMLFile           string   `json:"html_file"`
	NextDataFile       string   `json:"next_data_file"`
	BuildID            string   `json:"build_id"`
	NextPageDataURL    string   `json:"next_page_data_url,omitempty"`
	NextPageDataFile   string   `json:"next_page_data_file,omitempty"`
	JSONPagesBase      string   `json:"json_pages_base"`
	InitialListHeader  string   `json:"initial_list_header"`
	InitialListTag     string   `json:"initial_list_tag"`
	InitialMore        string   `json:"initial_more"`
	SavedPageFiles     []string `json:"saved_page_files"`
	CommanderCount     int      `json:"commander_count"`
	NormalizedJSONFile string   `json:"normalized_json_file,omitempty"`
	TableFile          string   `json:"table_file,omitempty"`
	Top200NamesFile    string   `json:"top_200_names_file,omitempty"`
}

// FetchResult bundles the normalized commander list together with the snapshot metadata.
type FetchResult struct {
	Commanders []TopCommander `json:"commanders"`
	Manifest   SnapshotManifest
}

// apiFetcher centralizes EDHREC HTTP calls so we can rate-limit and log them consistently.
type apiFetcher struct {
	client   *http.Client
	delay    time.Duration
	lastCall time.Time
}

// fetchURLBytes downloads one URL, waiting at least delay between calls and logging each request.
func (f *apiFetcher) fetchURLBytes(ctx context.Context, url string, pageLabel string) ([]byte, error) {
	if !f.lastCall.IsZero() {
		wait := time.Until(f.lastCall.Add(f.delay))
		if wait > 0 {
			timer := time.NewTimer(wait)
			defer timer.Stop()
			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("wait before %q canceled: %w", url, ctx.Err())
			case <-timer.C:
			}
		}
	}

	log.Printf("calling EDHREC API: %s url=%s", pageLabel, url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request for %q: %w", url, err)
	}
	req.Header.Set("User-Agent", "mtgconv-thank-you/1.0")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %q: %w", url, err)
	}
	defer resp.Body.Close()
	f.lastCall = time.Now()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %q: unexpected status %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", url, err)
	}

	return body, nil
}

// writeSnapshotFile creates parent directories as needed and writes one artifact to disk.
func writeSnapshotFile(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create directory for %q: %w", path, err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %q: %w", path, err)
	}
	return nil
}

// writeSnapshotJSON saves pretty-printed JSON for normalized outputs and the manifest.
func writeSnapshotJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal JSON for %q: %w", path, err)
	}
	return writeSnapshotFile(path, data)
}

// FetchTopCommanders retrieves the full ranked commander list currently shown on
// https://edhrec.com/commanders.
//
// What it does:
// 1. Download the HTML page
// 2. Extract the embedded __NEXT_DATA__ JSON bootstrap payload
// 3. Read the initial 100 commanders from cardlists[*].cardviews
// 4. Follow the "more" links at https://json.edhrec.com/pages/... until exhausted
// 5. Save raw payloads to disk as they are fetched
// 6. Return a single flat slice of commanders in rank order plus snapshot metadata
func FetchTopCommanders(ctx context.Context, client *http.Client, snapshotDir string) (*FetchResult, error) {
	const (
		commandersPageURL = "https://edhrec.com/commanders"
		jsonPagesBaseURL  = "https://json.edhrec.com/pages/"
	)
	fetcher := &apiFetcher{
		client: client,
		delay:  time.Second,
	}

	result := &FetchResult{
		Manifest: SnapshotManifest{
			FetchedAt:     time.Now().UTC().Format(time.RFC3339),
			SnapshotDir:   snapshotDir,
			CommandersURL: commandersPageURL,
			HTMLFile:      "commanders.html",
			NextDataFile:  "next-data.json",
			JSONPagesBase: jsonPagesBaseURL,
		},
	}

	// Step 1: fetch the public commanders page HTML and save the raw response body.
	htmlBytes, err := fetcher.fetchURLBytes(ctx, commandersPageURL, "bootstrap HTML")
	if err != nil {
		return nil, err
	}
	if err := writeSnapshotFile(filepath.Join(snapshotDir, result.Manifest.HTMLFile), htmlBytes); err != nil {
		return nil, err
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
	if err := writeSnapshotFile(filepath.Join(snapshotDir, result.Manifest.NextDataFile), []byte(nextDataJSON)); err != nil {
		return nil, err
	}

	var bootstrap nextData
	if err := json.Unmarshal([]byte(nextDataJSON), &bootstrap); err != nil {
		return nil, fmt.Errorf("decode __NEXT_DATA__: %w", err)
	}
	result.Manifest.BuildID = bootstrap.BuildID

	// Saving the current Next.js page-data response gives us an extra raw JSON copy
	// of the initial payload that is independent from the HTML wrapper.
	if bootstrap.BuildID != "" {
		result.Manifest.NextPageDataURL = fmt.Sprintf("https://edhrec.com/_next/data/%s/commanders.json", bootstrap.BuildID)
		result.Manifest.NextPageDataFile = "next-page-data.json"

		nextPageDataBytes, err := fetcher.fetchURLBytes(ctx, result.Manifest.NextPageDataURL, "page 0 next.js data")
		if err != nil {
			return nil, err
		}
		if err := writeSnapshotFile(filepath.Join(snapshotDir, result.Manifest.NextPageDataFile), nextPageDataBytes); err != nil {
			return nil, err
		}
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
	result.Manifest.InitialListHeader = current.Header
	result.Manifest.InitialListTag = current.Tag
	result.Manifest.InitialMore = current.More

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
	pageNumber := 1
	for nextMore != "" {
		pageURL := jsonPagesBaseURL + nextMore
		pageBytes, err := fetcher.fetchURLBytes(ctx, pageURL, fmt.Sprintf("page %d %s", pageNumber, nextMore))
		if err != nil {
			return nil, err
		}

		// Preserve the EDHREC relative page path under pages/... so the saved files
		// line up with the "more" chain we followed.
		pageRelativePath := filepath.ToSlash(filepath.Join("pages", filepath.FromSlash(nextMore)))
		if err := writeSnapshotFile(filepath.Join(snapshotDir, filepath.FromSlash(pageRelativePath)), pageBytes); err != nil {
			return nil, err
		}
		result.Manifest.SavedPageFiles = append(result.Manifest.SavedPageFiles, pageRelativePath)

		var page paginatedCommandersPage
		if err := json.Unmarshal(pageBytes, &page); err != nil {
			return nil, fmt.Errorf("decode paginated commanders page %q: %w", pageURL, err)
		}

		commanders = append(commanders, page.Cardviews...)

		if !page.IsPaginated || page.More == "" {
			break
		}
		nextMore = page.More
		pageNumber++
	}

	result.Commanders = commanders
	result.Manifest.CommanderCount = len(commanders)
	return result, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	// Save each run into a timestamped directory so past EDHREC snapshots remain available.
	snapshotDir := filepath.Join("edhrec", "snapshots", time.Now().UTC().Format("2006-01-02T15-04-05Z"))

	result, err := FetchTopCommanders(ctx, client, snapshotDir)
	if err != nil {
		panic(err)
	}
	commanders := result.Commanders

	// Build the full ranked table once so we can both print it and save it to disk.
	var table strings.Builder
	for _, c := range commanders {
		fmt.Fprintf(&table, "%3d. %-40s %6d decks\n", c.Rank, c.Name, c.NumDecks)
	}

	// Save a normalized JSON snapshot of the flattened commander list produced by this program.
	result.Manifest.NormalizedJSONFile = filepath.ToSlash(filepath.Join("normalized", "top-commanders.json"))
	if err := writeSnapshotJSON(filepath.Join(snapshotDir, filepath.FromSlash(result.Manifest.NormalizedJSONFile)), commanders); err != nil {
		panic(err)
	}

	// Save the ranked table as plain text.
	result.Manifest.TableFile = filepath.ToSlash(filepath.Join("normalized", "top-commanders-table.txt"))
	if err := writeSnapshotFile(filepath.Join(snapshotDir, filepath.FromSlash(result.Manifest.TableFile)), []byte(table.String())); err != nil {
		panic(err)
	}

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

	// Save the top-200 names as plain text.
	result.Manifest.Top200NamesFile = filepath.ToSlash(filepath.Join("normalized", "top-200-commander-names.txt"))
	if err := writeSnapshotFile(filepath.Join(snapshotDir, filepath.FromSlash(result.Manifest.Top200NamesFile)), []byte(names.String())); err != nil {
		panic(err)
	}

	// Save the manifest last so it can describe every artifact we wrote for this run.
	if err := writeSnapshotJSON(filepath.Join(snapshotDir, "manifest.json"), result.Manifest); err != nil {
		panic(err)
	}

	// Final output: summarize where the snapshot was written.
	fmt.Printf("saved EDHREC snapshot to %s\n", snapshotDir)
	fmt.Printf("saved %d commanders to %s\n", len(commanders), result.Manifest.TableFile)
	fmt.Printf("saved top %d names to %s\n", limit, result.Manifest.Top200NamesFile)

}
