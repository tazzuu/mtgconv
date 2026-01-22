package mtgconv

// implements an Interface which defines the methods which will be used to retrieve
// attributes from various API response objects
// in order to normalize them into canonical object types
//
// add more methods to these interfaces as we need more attributes in the canonical objects

type DeckReader interface {
	Meta() DeckMeta
	// Mainboard()
}


func NormalizeDeck(r DeckReader) Deck {
	return Deck{Meta: r.Meta()}
}