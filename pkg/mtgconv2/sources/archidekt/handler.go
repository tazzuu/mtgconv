package archidekt

import (
	"context"
	"net/http"
	"log/slog"
	// "strconv"

	"mtgconv/pkg/mtgconv2/core"
)

// input handler for decks from Moxfield

type Handler struct{}

func (h Handler) Source() core.APISource {
	return core.SourceArchidekt
}

func (h Handler) Fetch(ctx context.Context, input string, cfg core.Config, ovrr core.DeckMeta) (core.Deck, error) {
	_ = ctx
	_ = input
	_ = cfg

	slog.Debug("fetching Archidekt deck")
	deckID, err := DeckIDFromURL(input)
	if err != nil {
		return core.Deck{}, err
	}

	apiUrl := MakeAPIURL(deckID)
	slog.Debug("resolved API fetch URL", "url", apiUrl)

	// wait the required amount of time
	if err := APIRateLimiter.Wait(ctx); err != nil {
		return core.Deck{}, err
	}

	// start building the http request
	slog.Debug("building the http request")
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return core.Deck{}, err
	}
	req.Header.Set("Accept", "application/json")
	// slog.Debug("made http request", "req", req)

	// run the http request
	slog.Debug("running the http request")
	jsonStr, err := core.DoRequestJSON(req)
	if err != nil {
		return core.Deck{}, err
	}
	// slog.Debug("got jsonStr", "jsonStr", jsonStr)

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
	deckObj, err := MakeDeck(jsonStr)
	if err != nil {
		return core.Deck{}, err
	}

	// convert the Moxfield response deck type into the standardized core Deck type
	deck, err := DeckToCoreDeck(deckObj, input)
	if err != nil {
		return deck, err
	}

	return deck, nil
}

func (h Handler) Search(ctx context.Context, cfg core.Config, scfg core.SearchConfig) ([]core.DeckMeta, error) {
	_ = ctx
	_ = cfg

	return []core.DeckMeta{}, nil
}

func init() {
	core.RegisterSource(Handler{})
}
