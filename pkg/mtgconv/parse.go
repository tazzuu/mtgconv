package mtgconv

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

// get the final part of the URL
func DeckIDFromURL(rawUrl string) (string, error) {
	// rawUrl := "https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ"
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("could not get Deck ID from URL %v ; %v", rawUrl, err)
	}

	trimmed := strings.TrimSuffix(u.Path, "/")
	last := path.Base(trimmed)
	return last, nil
}
