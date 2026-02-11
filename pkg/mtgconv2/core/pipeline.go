package core

import (
	"context"
	"log/slog"
)

// runs the main data conversion pipeline using the input and output handlers selected
// TODO: move deckMetaOverride into an attribute on Config
func Run(ctx context.Context, cfg Config, deckMetaOverride DeckMeta) (string, Deck, error) {
	slog.Debug("starting processing pipeline")
	slog.Debug("configuring source handler")

	// initialize output variable
	var deck Deck

	// determine the correct input handler
	sourceHandler, err := HandlerForSource(cfg.InputSource)
	if err != nil {
		return "", Deck{}, err
	}

	// check the input type
	_, isFile := cfg.InputSource.Type()

	// if its a file, Import it
	if isFile {
		slog.Debug("fetching data from source")
		deck, err = sourceHandler.Import(cfg.UrlString, cfg)
		if err != nil {
			return "", Deck{}, err
		}
	} else {
		// its a URL so Fetch it
		slog.Debug("fetching data from source")
		deck, err = sourceHandler.Fetch(ctx, cfg.UrlString, cfg, deckMetaOverride)
		if err != nil {
			return "", Deck{}, err
		}
	}

	// get the output handler
	slog.Debug("configuring output handler")
	outputHandler, err := HandlerForOutput(cfg.OutputFormat)
	if err != nil {
		return "", Deck{}, err
	}

	// render the output
	rendered, err := outputHandler.Render(deck, cfg)
	if err != nil {
		return "", Deck{}, err
	}

	return rendered, deck, nil
}
