package core

import "context"

// interface definitions for the input and output handlers

type SourceHandler interface {
	Source() InputSource
	// method for retrieving a deck from an API endpoint
	Fetch(ctx context.Context, input string, cfg Config, ovrr DeckMeta) (Deck, error)
	// method for importing from a file
	Import(input string, cfg Config) (Deck, error)
	// method for performing Search on an API endpoint
	Search(ctx context.Context, cfg Config, scfg SearchConfig) ([]DeckMeta, error)
}

type OutputHandler interface {
	Format() OutputFormat
	Render(deck Deck, cfg Config) (string, error)
}
