package archidekt

import (
	"context"
	"net/http"
	"log/slog"
	"strconv"

	"mtgconv/pkg/mtgconv2/core"
)

// input handler for decks from Archidekt

type Handler struct{}

func (h Handler) Source() core.InputSource {
	return core.InputArchidektURL
}

func (h Handler) Import(filename string, cfg core.Config) (core.Deck, error) {
	_ = filename
	_ = cfg

	return core.Deck{}, nil
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

	// convert the Archidekt response deck type into the standardized core Deck type
	deck, err := DeckToCoreDeck(deckObj, input)
	if err != nil {
		return deck, err
	}

	return deck, nil
}

func (h Handler) Search(ctx context.Context, cfg core.Config, scfg core.SearchConfig) ([]core.DeckMeta, error) {
	_ = ctx
	_ = cfg
	slog.Debug("starting Archidekt Search")
	slog.Debug("Got search config", "scfg", scfg)

	var pageStart int = scfg.PageStart
	var pageEnd int = scfg.PageEnd
	// NOTE: Archidekt API does not have pageSize so limit the total number of output items instead
	var numDecks int = scfg.PageSize
	deckMetaList := []core.DeckMeta{}
	for page := pageStart; page <= pageEnd; page++ {
		// start building http request
		slog.Debug("building the http request")
		req, err := http.NewRequest(http.MethodGet, DeckSearchUrl, nil)
		if err != nil {
			return []core.DeckMeta{}, err
		}
		req.Header.Set("Accept", "application/json")

		// start appending query params
		q := req.URL.Query()
		q.Add("page", strconv.Itoa(page))

		// parse the arg for the sort type and order
		sortType, err := CoreSortTypeToArkSortType(scfg.SortType)
		if err != nil {
			return []core.DeckMeta{}, err
		}
		var sortArg string = sortType
		if scfg.SortDirection == core.SortDesc {
			sortArg = "-" + sortArg
		}
		q.Add("orderBy", sortArg)

		// parse the arg for the deck format
		formatInt, err := CoreFormatToDeckFormat(scfg.DeckFormat)
		if err != nil {
			return []core.DeckMeta{}, err
		}
		q.Add("deckFormat", strconv.Itoa(formatInt))

		q.Add("edhBracket", scfg.MinBracket.String())
		// q.Add("maxBracket", scfg.MaxBracket.String()) // NOTE: there is only one param for bracket in this Search

		if scfg.Username != "" {
			q.Add("ownerUsername", scfg.Username)
			// TODO: for username search also include these ; includePinned=true showIllegal=true board=mainboard
		}
		req.URL.RawQuery = q.Encode()

		slog.Debug("got query URL", "url", req.URL.String())

		// wait the required amount of time
		if err := APIRateLimiter.Wait(ctx); err != nil {
			return []core.DeckMeta{}, err
		}

		// run the http request
		slog.Debug("running the http request")
		jsonStr, err := core.DoRequestJSON(req)
		if err != nil {
			return []core.DeckMeta{}, err
		}

		// save JSON to file if that was requested
		if cfg.SaveJSON {
			// indent the JSON for readability
			pretty, err := core.PrettyJSON(jsonStr)
			if err != nil {
				return []core.DeckMeta{}, err
			}
			if err := core.SaveTxtToFile(core.ResponseJSONFilename, pretty); err != nil {
				return []core.DeckMeta{}, err
			}
		}

		// convert to Go object
		slog.Debug("converting to Go object")
		result, err := MakeSeachResult(jsonStr)
		if err != nil {
			slog.Error("error parsing JSON", "err", err)
			return []core.DeckMeta{}, err
		}

		slog.Debug("get results","page", page, "nresults", len(result.Results))

		// convert each search result to a core.DeckMeta
		for _, entry := range result.Results {
			if len(deckMetaList) >= numDecks {
				slog.Debug("Archidekt 'page size' limit reached, skipping the rest of the results", "numDecks", numDecks)
				break
			}
			deckMeta, err := SearchResultToDeckMeta(entry)
			if err != nil {
				return []core.DeckMeta{}, err
			}
			deckMetaList = append(deckMetaList, deckMeta)
		}
	}

	return deckMetaList, nil
}

func init() {
	core.RegisterSource(Handler{})
}
