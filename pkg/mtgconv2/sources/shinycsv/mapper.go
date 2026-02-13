package shinycsv

import (
	"log/slog"
	"mtgconv/pkg/mtgconv2/core"
	"mtgconv/pkg/mtgconv2/sets"
)

func ShinyRowsToCoreDeck(rows []*ShinyRow, config core.Config) (core.Deck) {
	// get the MTG Sets index
	setIndex := sets.GetSetIndex()

	// check some config parameters
	// TODO: get this from config instead of hard-coding it here!
	var includeTokens bool = false
	var includeArtCards bool = false

	deck := core.Deck{
		Meta: core.DeckMeta{
			Name: "Collection",
			Description: "Shiny app csv exported collection data",
		},
		// start empty sections map
		Sections: map[core.BoardType][]core.DeckEntry{},
	}

	for _, row := range rows {
		// check if card is valid before continuing
		// skip invalid cards e.g. Tokens, non-MTG, etc
		if !ValidateCard(row, includeTokens, includeArtCards) {
			slog.Debug("Skipping invalid card", "row", row)
			continue
		}
		// fix the names because they are TCG Player names with a bunch of junk added in them
		cleanName, _ := ParseCardName(row.ProductName)

		// check the name a second time after cleaning
		rowCopy := row
		rowCopy.ProductName = cleanName
		if !ValidateCard(rowCopy, includeTokens, includeArtCards) {
			slog.Debug("Skipping invalid card", "row", row)
			continue
		}

		var finishType core.FinishType = core.FinishDefault
		if row.IsFoil() {
			finishType = core.FinishFoil
		}

		// look up a proper set code if it exists
		// NOTE: some of the rows already use set codes and others use set names
		// check first if the Name is actually a code
		// var setName string = row.SetName
		// var setCode string
		// var setType string

		// validate set codes
		var setName string
		var setCode string
		var setType string
		setName, setCode, setType = ValidateSet(row, setIndex)

		entry := core.DeckEntry{
			Quantity: row.Quantity,
			Board: core.BoardMain,
			Finish: finishType,
			Card: core.Card{
				Name: cleanName,
				CollectorNumber: CleanCollectorNumber(row.Discriminator),
				Tags: []string{row.Tag},
				SetName: setName,
				SetCode: setCode,
				SetType: setType,
				GroupName: row.GroupName,
				Rarity: row.Rarity,
				DateAdded: row.DateAdded,
				IsProxy: row.IsProxy(),
			},
		}

		err := deck.AddToSection(core.BoardMain, entry)

		if err != nil {
			slog.Error("error adding entry to Deck", "err", err)
		}
	}

	return deck
}