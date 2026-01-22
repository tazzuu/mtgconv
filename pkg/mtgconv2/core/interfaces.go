package core

import "context"

type SourceHandler interface {
	Source() APISource
	Fetch(ctx context.Context, input string, cfg Config) (Deck, error)
}

type OutputHandler interface {
	Format() OutputFormat
	Render(deck Deck, cfg Config) (string, error)
}
