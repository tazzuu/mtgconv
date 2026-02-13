package moxfieldcollection

import (
	"strings"
	"strconv"
	"mtgconv/pkg/mtgconv2/core"
)

// convert the core object to the output object
func CoreDeckEntryToMoxfieldEntry(entry core.DeckEntry) MoxfieldCollectionRow {
	var foilStr string
	if entry.Finish == core.FinishFoil || entry.Finish == core.FinishEtched {
		foilStr = "foil"
	}
	return MoxfieldCollectionRow{
		Count: entry.Quantity,
		TradelistCount: entry.Quantity,
		Name: entry.Card.Name,
		Edition: entry.Card.SetCode, // NOTE: also see entry.Card.SetName
		// Condition: "",
		Language: "English",
		Foil: foilStr,
		// Tags: ,
		LastModified: entry.Card.DateAdded,
		CollectorNumber: entry.Card.CollectorNumber,
		// Alter:
		Proxy: strings.ToUpper(strconv.FormatBool(entry.Card.IsProxy)),
		PlaytestCard: strings.ToUpper(strconv.FormatBool(entry.Card.IsProxy)),
		// PurchasePrice:
	}
}