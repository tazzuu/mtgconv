package archidekt

import (
	"log/slog"
	"strings"
	"mtgconv/pkg/mtgconv2/core"
	"strconv"
	"net/url"
)

func DeckToCoreDeck(ark DeckResponse, url string) (core.Deck, error) {
	//
	// NOTE: references ; https://github.com/lheyberger/mtg-parser/blob/master/src/mtg_parser/archidekt.py ; https://github.com/linkian209/pyrchidekt/blob/main/pyrchidekt/deck.py
	//

	// get the deck format value and convert to internal DeckFormat
	slog.Debug("got deck", "ark.Name", ark.Name, "ark.ID", ark.ID, "ark.DeckFormat", ark.DeckFormat)
	deckFormat, err := DeckFormatToCoreFormat(ark.DeckFormat)
	if err != nil {
		return core.Deck{}, err
	}

	// get the card "categories" included in the Deck
	categories := make(map[string]bool) // _, ok := m["route"]
	for _, val := range ark.Categories {
		if val.IncludedInDeck {
			categories[val.Name] = true
		}
	}


	// initialize core Deck object
	deck := core.Deck{
		Meta: core.DeckMeta{
			ID: strconv.Itoa(ark.ID),
			Name: ark.Name,
			Description: ark.Description,
			Format: string(deckFormat),
			URL: url,
			Authors: []string{ark.Owner.Username},
			Source: ApiSource,
			Date: ark.RetrievedAt,
			CreatedAt: ark.CreatedAt.Format("2006-01-02"),
			UpdatedAt: ark.UpdatedAt.Format("2006-01-02"),
			// Version: ark.Version,
		},
		// start empty sections map
		Sections: make(map[core.BoardType][]core.DeckEntry),
	}

	// add each card to the correct deck section based on the categories
	for _, card := range ark.Cards {
		isInDeck, isCommander := CategoryIsInDeck(card.Categories, categories)
		if isCommander {
			slog.Debug("got Commander", "card.Name", card.Card.OracleCard.Name)
			entry := DeckEntryToCoreDeckEntry(card, core.BoardCommander)
			err := deck.AddToSection(core.BoardCommander, entry)
			if err != nil {
				return core.Deck{}, err
			}
		} else if isInDeck {
			slog.Debug("adding to deck", "card.Name", card.Card.OracleCard.Name)
			entry := DeckEntryToCoreDeckEntry(card, core.BoardMain)
			err := deck.AddToSection(core.BoardMain, entry)
			if err != nil {
				return core.Deck{}, err
			}
		}
	}

	return deck, nil
}

func DeckEntryToCoreDeckEntry(entry DeckCardEntry, boardType core.BoardType) core.DeckEntry {
	return core.DeckEntry{
		Quantity: entry.Quantity,
			Board: boardType,
			Finish: core.FinishType(core.FinishDefault), // TODO: use the default Finish for all cards because I dont know if Archidekt lists the Foil status
			Card: core.Card{
				Name: entry.Card.OracleCard.Name,
				SetCode: entry.Card.Edition.EditionCode, // TODO: figure out if this is the full name or the short code for the name
				SetType: entry.Card.Edition.EditionType,
				CollectorNumber: entry.Card.CollectorNumber,
				Layout: entry.Card.OracleCard.Layout,
				ManaCost: entry.Card.OracleCard.ManaCost,
				TypeLine: strings.Join(entry.Card.OracleCard.Types, ","),
				OracleText: entry.Card.OracleCard.Text,
				Power: entry.Card.OracleCard.Power,
				Toughness: entry.Card.OracleCard.Toughness,
				Colors: entry.Card.OracleCard.Colors,
				ColorIdentity: entry.Card.OracleCard.ColorIdentity,
				Rarity: entry.Card.Rarity,
				Language: entry.Card.OracleCard.Lang,
				IDs: core.CardIDs{
					// ScryfallID: entry.Card.ScryfallID,
					TCGPlayerID: entry.Card.TcgProductID,
					CardKingdomID: entry.Card.CkNormalID,
					MultiverseIDs: []int{entry.Card.MultiverseID},
				},
				NumFaces: len(entry.Card.OracleCard.Faces),
			},
	}
}

func SearchResultToDeckMeta(ark DeckSearchItem) (core.DeckMeta, error) {
	deckFormat, err := DeckFormatToCoreFormat(ark.DeckFormat)
	if err != nil {
		return core.DeckMeta{}, err
	}
	deckUrl, err := url.JoinPath(DeckPublicUrlBase, strconv.Itoa(ark.ID))
	if err != nil {
		return core.DeckMeta{}, err
	}
	meta := core.DeckMeta{
			ID: strconv.Itoa(ark.ID),
			Name: ark.Name,
			Format: string(deckFormat),
			URL: deckUrl,
			Date: ark.RetrievedAt,
			CreatedAt: ark.CreatedAt.Format("2006-01-02"),
			UpdatedAt: ark.UpdatedAt.Format("2006-01-02"),
			Authors: []string{ark.Owner.Username},
			Bracket: core.CommanderBracket(ark.EdhBracket),
			ViewCount: ark.ViewCount,
	}
	return meta, nil
}