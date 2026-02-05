package archidekt

import (
	"log/slog"
	"net/url"
	"strings"
	"path"
	"strconv"

	"mtgconv/pkg/mtgconv2/core"
)

// get the deck ID from the URL provided by the user
// https://archidekt.com/decks/4798129/kibo_uktabi_prince_mh3 -> https://archidekt.com/api/decks/4798129/
func DeckIDFromURL(rawUrl string) (string, error) {
	slog.Debug("getting the deck ID from the provided url", "url", rawUrl)

	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", &core.DeckIDParseError{URL:rawUrl}
	}

	// find the part of the path that says "decks" then take the part that comes after it
	segments := strings.Split(strings.Trim(u.Path, "/"), "/")
	for i, seg := range segments {
		if seg == "decks" && i+1 < len(segments) {
			return segments[i+1], nil
		}
	}

	return "", &core.DeckIDParseError{URL:rawUrl}
}

func MakeAPIURL(deckID string) string {
	u, _ := url.Parse(DeckFetchUrl)
	u.Path = path.Join(u.Path, deckID)
	return (u.String() + "/")
}

// https://github.com/linkian209/pyrchidekt/blob/main/pyrchidekt/formats.py
func DeckFormatToCoreFormat(i int) (core.DeckFormat, error) { // DeckFormatCommander
	switch i {
	case 3:
		return core.DeckFormatCommander, nil
	default:
		return "", &core.UnknownDeckFormat{Format:strconv.Itoa(i)}
	}
}

func CategoryIsInDeck(categories []string, catMap map[string]bool) (bool, bool) {
	var isInDeck bool
	var isCommander bool
	for _, cat := range categories {
		if _, exists := catMap[cat]; exists {
			isInDeck = true
			break
		}
	}
	for _, cat := range categories {
		if strings.ToLower(cat) == "commander" {
			isCommander = true
			break
		}
	}
	return isInDeck, isCommander
}