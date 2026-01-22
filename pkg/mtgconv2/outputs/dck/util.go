package dck

import (
	"fmt"
	"strings"
	"mtgconv/pkg/mtgconv2/core"
)

// apply conditional formatting for the card line in the .dck file
// enforce some compatibility requirements here
func FormatDckLine(entry core.DeckEntry, compatibilityMode bool) string {
	// if the card is from a Promo set, need to exclude the Set Code and Collector Number in the output
	if entry.Card.SetType == "promo" && compatibilityMode {
		return fmt.Sprintf("%d %s",
			entry.Quantity,
			entry.Card.Name,
		)
	}

	// default format
	return fmt.Sprintf("%d %s|%s|%s",
		entry.Quantity,
		entry.Card.Name,
		strings.ToUpper(entry.Card.SetCode),
		entry.Card.CollectorNumber,
	)
}