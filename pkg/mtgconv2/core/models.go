package core

import "time"

// top level deck object type
type Deck struct {
	Meta     DeckMeta `json:"meta"`
	Sections map[BoardType][]DeckEntry `json:"sections"` // Mainboard, Sideboard, etc.
}

// method to add an entry to a board, with validations
func (d *Deck) AddToSection(section BoardType, entry DeckEntry) error {
	// make sure its a recognized board type being passed
	boardType, err := ParseBoardType(string(section))
	if err != nil {
		return err
	}
	// make sure quantity is >0
	if entry.Quantity <= 0 {
		return &InvalidQuantity{entry.Quantity}
	}
	// make sure its a recognized card finish
	_, err = ParseFinishType(string(entry.Finish))
	if err != nil {
		return err
	}
	// make sure the section is initialized
	if d.Sections == nil {
		d.Sections = map[BoardType][]DeckEntry{}
	}
	// do more validations here
	d.Sections[boardType] = append(d.Sections[boardType], entry)
	return nil
}

// metadata for a deck
type DeckMeta struct {
	ID          string `json:"id"`// alphanumeric short id for the deck
	PublicID string `json:"publicId"`// longer ID used by some API's such as Moxfield
	Name        string `json:"name"`
	Description string `json:"description"`
	Format      string `json:"format"`
	Visibility  string `json:"visibility"`
	URL         string `json:"url"`
	Date        time.Time `json:"date"`// date retrieved from API
	CreatedAt string `json:"createdAt"`// time in UTC
	UpdatedAt string `json:"updatedAt"`// time in UTC
	Authors     []string `json:"authors"`
	Bracket CommanderBracket `json:"bracket"`
	LikeCount int `json:"likeCount"`
	ViewCount int `json:"viewCount"`
	BookmarkCount int `json:"bookmarkCount"`
	CommentCount int `json:"commentCount"`
	Version int `json:"version"`
	Source      APISource `json:"source"`
}

// an entry for one or more copies of a specific card in the Mainboard, Sideboard, etc.
type DeckEntry struct {
	Quantity int `json:"quantity"`
	Board    BoardType `json:"board"`
	Finish   FinishType `json:"finish"`
	Card     Card `json:"card"`
}

type Card struct {
	Name            string `json:"name"`
	SetCode         string `json:"setCode"`
	SetType string `json:"setType"`
	CollectorNumber string `json:"collectorNumber"`
	Layout          string `json:"layout"`
	ManaCost        string `json:"manaCost"`
	TypeLine        string `json:"typeLine"`
	OracleText      string `json:"oracleText"`
	Power           string `json:"power"`
	Toughness       string `json:"toughness"`
	Colors          []string `json:"colors"`
	ColorIdentity   []string `json:"colorIdentity"`
	Rarity          string `json:"rarity"`
	Language        string `json:"language"`
	IDs             CardIDs `json:"ids"`
	NumFaces int `json:"numFaces"`
}

type CardIDs struct {
	ScryfallID    string `json:"scryfallId"`
	MultiverseIDs []int `json:"multiversIds"`
	TCGPlayerID   int `json:"tcgPlayerId"`
	CardKingdomID int `json:"cardKingdomId"`
}

