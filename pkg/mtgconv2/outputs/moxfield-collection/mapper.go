package moxfieldcollection

import (
"mtgconv/pkg/mtgconv2/core"
)

func CoreDeckEntryToMoxfieldEntry(entry core.DeckEntry) MoxfieldCollectionRow {
	return MoxfieldCollectionRow{
		Count: entry.Quantity,
		TradelistCount: entry.Quantity,
		Name: entry.Card.Name,
		Edition: entry.Card.SetCode, // NOTE: also see entry.Card.SetName
		// Condition: "",
		Language: "English",
		// Foil: "",
		// Tags: ,
		LastModified: entry.Card.DateAdded,
		CollectorNumber: entry.Card.CollectorNumber,
		// Alter:
		// Proxy:
		// PlaytestCard:
		// PurchasePrice:
	}
}