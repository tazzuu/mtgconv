package mtgconv

type Deck struct {
	Meta DeckMeta
}

type DeckMeta struct {
	ID                 string
	Name               string
	Description        string
	Format             string
	Visibility         string
	URL          string
	Date string
	Authors          []string
}
