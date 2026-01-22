package mtgconv

import (
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strconv"
	"strings"

	_ "embed"
	"text/template"
)

//go:embed templates/dck.txt
var dckTemplateStr string

func MoxfieldURLtoDckFormat(config Config) (string, error) {
	var result string

	jsonStr, err := FetchMoxfieldDecklistJson(config)

	// convert the JSON string into Go objects
	deck, err := MakeMoxfieldDeck(jsonStr)
	if err != nil {
		return result, fmt.Errorf("error while parsing Moxfield API response: %v", err)
	}

	// convert to final .dck decklist format
	result, err = MoxfieldDeckToDckFormat(deck)
	if err != nil {
		return result, fmt.Errorf("error while generating the .dck template: %v", err)
	}
	return result, nil
}

// convert the Moxfield JSON response object into a .dck decklist format
func MoxfieldDeckToDckFormat(deck MoxfieldDeck) (string, error) {
	var result string
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"NumericOrZero": func(s string) string {
			trimmed := strings.TrimSpace(s)
			if _, err := strconv.Atoi(trimmed); err != nil {
				return "0"
			}
			return trimmed
		},
	}
	tmpl, err := template.New("dck").Funcs(funcMap).Parse(dckTemplateStr)
	if err != nil {
		slog.Error("Error initializing report template", "err", err)
		return "", fmt.Errorf("initializing report template: %v", err)
	}

	var output strings.Builder
	if err := tmpl.Execute(&output, deck); err != nil {
		slog.Error("Error creating report", "err", err)
		return "", fmt.Errorf("creating report: %v", err)
	}

	result = output.String()

	return result, nil
}










//
//
//
// OLD method of converting to .dck text format
//
func MoxfieldDeckToDckFormatOLD(deck MoxfieldDeck) string {
	var result string = ""
	var lines []string

	// make the metadata section
	// TODO: replace this with template
	lines = append(lines, []string{
		"[metadata]",
		fmt.Sprintf("Name=%s", deck.Name),
	}...)

	// make the Main section
	lines = append(lines, "[Main]")

	if len(deck.Mainboard) > 0 {
		for key, value := range deck.Mainboard {
			lines = append(lines, fmt.Sprintf(
				"%d %v|%v|[%v]",
				value.Quantity,
				key,
				strings.ToUpper(value.Card.Set),
				value.Card.CN,
			))
		}
	}

	// make the Commander section
	if len(deck.Commanders) > 0 {
		lines = append(lines, "[Commander]")
		for key, value := range deck.Commanders {
			lines = append(lines, fmt.Sprintf(
				"%d %v|%v|[%v]",
				value.Quantity,
				key,
				strings.ToUpper(value.Card.Set),
				value.Card.CN,
			))
		}
	}

	// add Sideboard section
	if len(deck.Sideboard) > 0 {
		lines = append(lines, "[Sideboard]")
		// dont try to add more items than the limit or the sideboard size
		sideboardSize := len(deck.Sideboard)
		sideboardLimit := sideboardMaxSize
		if sideboardSize <= sideboardLimit {
			sideboardLimit = sideboardSize
		}
		// get the sideboard card keys in order
		sideboardKeys := slices.Collect(maps.Keys(deck.Sideboard))
		slices.Sort(sideboardKeys)
		// add only the number of items up to the limit
		for i := 0; i < sideboardLimit; i++ {
			card := deck.Sideboard[sideboardKeys[i]]
			lines = append(lines, fmt.Sprintf(
				"%d %v|%v|[%v]",
				card.Quantity,
				sideboardKeys[i],
				strings.ToUpper(card.Card.Set),
				card.Card.CN,
			))
		}
	}

	result = strings.Join(lines, "\n")

	return result
}
