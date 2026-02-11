package shinycsv

import (
	"log/slog"
	"mtgconv/pkg/mtgconv2/core"
)

func ShinyRowsToCoreDeck(rows []*ShinyRow) (core.Deck) {
	deck := core.Deck{
		Meta: core.DeckMeta{
			Name: "Collection",
			Description: "Shiny app csv exported collection data",
		},
		// start empty sections map
		Sections: map[core.BoardType][]core.DeckEntry{},
	}

	for _, row := range rows {
		cleanName, _ := ParseCardName(row.ProductName)

		var finishType core.FinishType = core.FinishDefault
		if row.IsFoil() {
			finishType = core.FinishFoil
		}

		entry := core.DeckEntry{
			Quantity: row.Quantity,
			Board: core.BoardMain,
			Finish: finishType,
			Card: core.Card{
				Name: cleanName,
				CollectorNumber: CleanCollectorNumber(row.Discriminator),
				Tags: []string{row.Tag},
				SetName: row.SetName,
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