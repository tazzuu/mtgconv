package mtgconv

import (
	"net/url"
	"path"
	"strings"
)

// get the final part of the URL
func DeckIDFromURL(rawUrl string) string {
	// rawUrl := "https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ"
	u, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	trimmed := strings.TrimSuffix(u.Path, "/")
	last := path.Base(trimmed)
	return(last) // Wrcumkgcc0qjIB2bwoDvqQ
}
