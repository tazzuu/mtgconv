package moxfield

import (
	"log/slog"
	"net/url"
	"strings"
	"path"

	"mtgconv/pkg/mtgconv2/core"
)

func MakeMoxfieldAPIUrl(deckID string) string {
	u, _ := url.Parse(MoxfieldBaseUrl)
	u.Path = path.Join(u.Path, deckID)
	return (u.String())
}

// get the final part of the URL
func DeckIDFromURL(rawUrl string) (string, error) {
	// rawUrl := "https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ"
	slog.Debug("getting the deck ID from the provided url", "url", rawUrl)

	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", &core.DeckIDParseError{URL:rawUrl}
	}

	trimmed := strings.TrimSuffix(u.Path, "/")
	last := path.Base(trimmed)

	slog.Debug("got deck ID", "deckID", last)
	return last, nil
}

