package moxfield

import (
	"context"
	"net/http"
	"log/slog"

	"mtgconv/pkg/mtgconv2/core"
)

type Handler struct{}

func (h Handler) Source() core.APISource {
	return core.SourceMoxfield
}

func (h Handler) Fetch(ctx context.Context, input string, cfg core.Config) (core.Deck, error) {
	_ = ctx
	_ = input
	_ = cfg

	// get the deck ID from the provided URL
	deckID, err := DeckIDFromURL(input)
	if err != nil {
		return core.Deck{}, err
	}

	// create the API query URL
	slog.Debug("making API query URL")
	deckAPIUrl := MakeMoxfieldAPIUrl(deckID)
	slog.Debug("Got API Query url", "deckAPIUrl", deckAPIUrl)

	// wait the required amount of time
	if err := MoxfieldAPIRateLimiter.Wait(ctx); err != nil {
		return core.Deck{}, err
	}

	// start building the http request
	slog.Debug("building the http request")
	req, err := http.NewRequest(http.MethodGet, deckAPIUrl, nil)
	if err != nil {
		return core.Deck{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", cfg.UserAgent)

	// run the http request
	slog.Debug("running the http request")
	jsonStr, err := core.DoRequestJSON(req)
	if err != nil {
		return core.Deck{}, err
	}

	// convert the JSON string into Go objects
	slog.Debug("parsing the http request JSON")
	moxfieldDeck, err := MakeMoxfieldDeck(jsonStr)
	if err != nil {
		return core.Deck{}, err
	}

	// convert the Moxfield response deck type into the standardized core Deck type
	deck, err := MoxfieldDeckToCoreDeck(moxfieldDeck)
	if err != nil {
		return deck, err
	}

	return deck, nil
	// return core.Deck{}, fmt.Errorf("moxfield source handler not implemented")
}

func init() {
	core.RegisterSource(Handler{})
}
