package moxfield

import (
	"log/slog"
	"mtgconv/pkg/mtgconv2/core"
)

func MoxfieldEntryToCoreEntry(mxEntry MoxfieldDeckEntry, boardType core.BoardType) (core.DeckEntry) {
	return core.DeckEntry{
			Quantity: mxEntry.Quantity,
			Board: boardType,
			Finish: core.FinishType(mxEntry.Finish),
			Card: core.Card{
				Name: mxEntry.Card.Name, // TODO: figure out of the Card.Name always matches the cardName here
				SetCode: mxEntry.Card.Set, // TODO: figure out if this is the full name or the short code for the name
				CollectorNumber: mxEntry.Card.CN,
				Layout: mxEntry.Card.Layout,
				ManaCost: mxEntry.Card.ManaCost,
				TypeLine: mxEntry.Card.TypeLine,
				OracleText: mxEntry.Card.OracleText,
				Power: mxEntry.Card.Power,
				Toughness: mxEntry.Card.Toughness,
				Colors: mxEntry.Card.Colors,
				ColorIdentity: mxEntry.Card.ColorIdentity,
				Rarity: mxEntry.Card.Rarity,
				Language: mxEntry.Card.Lang,
				IDs: core.CardIDs{
					ScryfallID: mxEntry.Card.ScryfallID,
					TCGPlayerID: mxEntry.Card.TCGPlayerID,
					CardKingdomID: mxEntry.Card.CardKingdomID,
					MultiverseIDs: mxEntry.Card.MultiverseIDs,
				},
			},
		}
}

func MoxfieldDeckToCoreDeck(mx MoxfieldDeck) (core.Deck, error) {
	// get the authors list from the response object
	authors := []string{}
	for _, author := range mx.Authors {
		authors = append(authors, author.UserName)
	}

	// initialize core Deck object
	deck := core.Deck{
		Meta: core.DeckMeta{
			ID: mx.ID,
			Name: mx.Name,
			Description: mx.Description,
			Format: mx.Format,
			Visibility: mx.Visibility,
			URL: mx.PublicURL,
			Authors: authors,
			Source: ApiSource,
			Date: mx.RetrievedAt,
			CreatedAt: mx.CreatedAtUTC,
			UpdatedAt: mx.LastUpdatedAtUTC,
			Version: mx.Version,
		},
		// start empty sections map
		Sections: map[core.BoardType][]core.DeckEntry{},
	}

	// add each of the sections from the Moxfield response object
	// start with Mainboard
	for cardName, mxEntry := range mx.Mainboard {
		slog.Debug("parsing card", "cardName", cardName)
		entry := MoxfieldEntryToCoreEntry(mxEntry, core.BoardMain)
		err := deck.AddToSection(core.BoardMain, entry)
		if err != nil {
			// TODO: how should we handle errors here??
			slog.Error("error adding entry to Deck", "err", err)
		}
	}

	// add sideboard
	for cardName, mxEntry := range mx.Sideboard {
		slog.Debug("parsing card", "cardName", cardName)
		entry := MoxfieldEntryToCoreEntry(mxEntry, core.BoardSideboard)
		err := deck.AddToSection(core.BoardSideboard, entry)
		if err != nil {
			// TODO: how should we handle errors here??
			slog.Error("error adding entry to Deck", "err", err)
		}
	}

	// add commanders
	for cardName, mxEntry := range mx.Commanders {
		slog.Debug("parsing card", "cardName", cardName)
		entry := MoxfieldEntryToCoreEntry(mxEntry, core.BoardCommander)
		err := deck.AddToSection(core.BoardCommander, entry)
		if err != nil {
			// TODO: how should we handle errors here??
			slog.Error("error adding entry to Deck", "err", err)
		}
	}

	return deck, nil
}