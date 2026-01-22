package mtgconv

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"log/slog"
	"context"
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
		return "", fmt.Errorf("could not get Deck ID from URL %v ; %v", rawUrl, err)
	}

	trimmed := strings.TrimSuffix(u.Path, "/")
	last := path.Base(trimmed)

	slog.Debug("got deck ID", "deckID", last)
	return last, nil
}


func FetchMoxfieldDecklistJson(config Config) (string, error) {
	var result string

	// get the deck ID from the provided URL
	deckID, err := DeckIDFromURL(config.UrlString)
	if err != nil {
		return result, fmt.Errorf("error getting deck ID from URL: %v", err)
	}

	// create the API query URL
	slog.Debug("making API query URL")
	deckAPIUrl := MakeMoxfieldAPIUrl(deckID)
	slog.Debug("Got API Query url", "deckAPIUrl", deckAPIUrl)

	// fetch the JSON query result
	// TODO: add context here
	jsonStr, err := RequestJSON(context.TODO(), deckAPIUrl, config.UserAgent)
	if err != nil {
		return result, fmt.Errorf("error while getting the API query result: %v", err)
	}
	slog.Debug("got API query result")

	return jsonStr, nil
}
