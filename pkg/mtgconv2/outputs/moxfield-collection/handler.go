package moxfieldcollection

import (
	"github.com/gocarina/gocsv"
	"mtgconv/pkg/mtgconv2/core"
	"bytes"
)

type Handler struct{}

func (h Handler) Format() core.OutputFormat {
	return core.OutputMoxfieldCollection
}

// https://moxfield.com/help/importing-collection
func (h Handler) Render(deck core.Deck, cfg core.Config) (string, error) {
	entries := []core.DeckEntry{}
	for _, boardType := range core.BoardTypes() {
		entries = append(entries, deck.Sections[boardType]...)
	}
	collectionEntries := []MoxfieldCollectionRow{}
	for _, entry := range entries {
		collectionEntries = append(collectionEntries, CoreDeckEntryToMoxfieldEntry(entry))
	}

	var buf bytes.Buffer

	if err := gocsv.Marshal(&collectionEntries, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func init() {
	core.RegisterOutput(Handler{})
}
