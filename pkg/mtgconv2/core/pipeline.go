package core

import (
	"context"
	"log/slog"
)

// runs the main data conversion pipeline using the input and output handlers selected
func Run(ctx context.Context, cfg Config) (string, error) {
	slog.Debug("starting processing pipeline")
	src, err := DetectURLSource(cfg.UrlString)
	if err != nil {
		return "", err
	}

	slog.Debug("configuring source handler")
	sourceHandler, err := HandlerForSource(src)
	if err != nil {
		return "", err
	}

	slog.Debug("fetching data from source")
	deck, err := sourceHandler.Fetch(ctx, cfg.UrlString, cfg)
	if err != nil {
		return "", err
	}

	slog.Debug("configuring output handler")
	outputHandler, err := HandlerForOutput(cfg.OutputFormat)
	if err != nil {
		return "", err
	}

	return outputHandler.Render(deck, cfg)
}
