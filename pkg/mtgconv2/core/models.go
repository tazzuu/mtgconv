package core

import "time"

// top level deck object type
type Deck struct {
	Meta     DeckMeta
	Sections map[BoardType][]DeckEntry // Mainboard, Sideboard, etc.
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
	ID          string // alphanumeric short id for the deck
	PublicID string // longer ID used by some API's such as Moxfield
	Name        string
	Description string
	Format      string
	Visibility  string
	URL         string
	Date        time.Time // date retrieved from API
	CreatedAt string // time in UTC
	UpdatedAt string // time in UTC
	Authors     []string
	Bracket CommanderBracket
	LikeCount int
	ViewCount int
	CommentCount int
	Version int
	Source      APISource
}

// an entry for one or more copies of a specific card in the Mainboard, Sideboard, etc.
type DeckEntry struct {
	Quantity int
	Board    BoardType
	Finish   FinishType
	Card     Card
}

type Card struct {
	Name            string
	SetCode         string
	SetType string
	CollectorNumber string
	Layout          string
	ManaCost        string
	TypeLine        string
	OracleText      string
	Power           string
	Toughness       string
	Colors          []string
	ColorIdentity   []string
	Rarity          string
	Language        string
	IDs             CardIDs
	NumFaces int
}

type CardIDs struct {
	ScryfallID    string
	MultiverseIDs []int
	TCGPlayerID   int
	CardKingdomID int
}
