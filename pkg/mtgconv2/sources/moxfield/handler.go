package moxfield

import (
	"context"
	"net/http"
	"log/slog"

	"mtgconv/pkg/mtgconv2/core"
)

// input handler for decks from Moxfield

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

	// save JSON to file if that was requested
	if cfg.SaveJSON {
		// indent the JSON for readability
		pretty, err := core.PrettyJSON(jsonStr)
		if err != nil {
			return core.Deck{}, err
		}
		if err := core.SaveTxtToFile(core.ResponseJSONFilename, pretty); err != nil {
			return core.Deck{}, err
		}
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

func (h Handler) Search(ctx context.Context, cfg core.Config, scfg core.SearchConfig) (string, error) {
	_ = ctx
	_ = cfg
	slog.Debug("starting Moxfield Search")
	slog.Debug("Got search config", "scfg", scfg)

	// start building http request
	slog.Debug("building the http request")
	req, err := http.NewRequest(http.MethodGet, MoxfieldDeckSearchUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", cfg.UserAgent)

	// start appending query params
	q := req.URL.Query()
	q.Add("pageNumber", "1")
	q.Add("pageSize", "64")
	q.Add("sortType", string(scfg.SortType))
	q.Add("sortDirection", string(scfg.SortDirection))
	q.Add("fmt", string(scfg.DeckFormat))
	q.Add("minBracket", scfg.MinBracket.String())
	q.Add("maxBracket", scfg.MaxBracket.String())
	req.URL.RawQuery = q.Encode()

	slog.Debug("got query URL", "url", req.URL.String())

	// run the http request
	slog.Debug("running the http request")
	jsonStr, err := core.DoRequestJSON(req)
	if err != nil {
		return "", err
	}

	return jsonStr, nil
}

func init() {
	core.RegisterSource(Handler{})
}
