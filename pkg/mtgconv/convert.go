package mtgconv

import (
	"fmt"
	"strings"
	"maps"
	"slices"
)

func MoxfieldDeckToDckFormat(deck DeckResponse) string {
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
