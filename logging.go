package main

import (
	"log/slog"
	"os"
)

// initialize internal logger
func configureLogging (verbose bool) {

	// default log level is INFO
	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	// create a TextHandler and set it globally
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level:     level,
		AddSource: false, // adds the line number from source code
	})
	slog.SetDefault(slog.New(handler))
	slog.Debug("logger initialized", "verbose", verbose)
}