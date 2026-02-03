package core

import "context"

// interface definitions for the input and output handlers

type SourceHandler interface {
	Source() APISource
	Fetch(ctx context.Context, input string, cfg Config, ovrr DeckMeta) (Deck, error)
	// TODO: implement Search function here ; change this output type to []Deck
	Search(ctx context.Context, cfg Config, scfg SearchConfig) ([]DeckMeta, error)
}

type OutputHandler interface {
	Format() OutputFormat
	Render(deck Deck, cfg Config) (string, error)
}
