package dck

import (
	"log/slog"
	"fmt"
	"strings"
	"mtgconv/pkg/mtgconv2/core"
)

// apply conditional formatting for the card line in the .dck file
// enforce some compatibility requirements here
func FormatDckLine(entry core.DeckEntry, compatibilityMode bool) string {
	quantity := entry.Quantity
	cardName := entry.Card.Name
	setCode := strings.ToUpper(entry.Card.SetCode)
	collectorNumber := entry.Card.CollectorNumber
	setType := entry.Card.SetType
	numFaces := entry.Card.NumFaces

	if compatibilityMode {
		// if the card has multiple faces and has // in the name, return only the first face e.g. first part of the name
		if strings.Contains(cardName, "//") {
			slog.Debug("enforcing compat name for card with // in the name", "name", cardName)
			parts := core.SplitMultiFaceName(cardName)
			cardName = parts[0]
			slog.Debug("new name;", "name", cardName)
		}

		if numFaces > 1 {
			slog.Debug("enforcing compat name for card with multiple faces", "name", cardName)
			parts := core.SplitMultiFaceName(cardName)
			cardName = parts[0]
			slog.Debug("new name;", "name", cardName)
		}

		// if the card is from a Promo set, need to exclude the Set Code and Collector Number in the output
		// the PLST set also causes issues too
		if setType == "promo" || setCode == "PLST"  {
			slog.Debug("enforcing compat set codes", "name", cardName, "setType", setType, "setCode", setCode)
			return fmt.Sprintf("%d %s",
				quantity,
				cardName,
			)
		}

	}


	// default format
	return fmt.Sprintf("%d %s|%s|%s",
		quantity,
		cardName,
		setCode,
		collectorNumber,
	)
}