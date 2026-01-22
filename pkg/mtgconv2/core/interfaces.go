package core

import "context"

// interface definitions for the input and output handlers

type SourceHandler interface {
	Source() APISource
	Fetch(ctx context.Context, input string, cfg Config) (Deck, error)
}

type OutputHandler interface {
	Format() OutputFormat
	Render(deck Deck, cfg Config) (string, error)
}
