package shinycsv

import (
	"context"
	// "net/http"
	// "log/slog"
	// "strconv"

	"mtgconv/pkg/mtgconv2/core"
)

// input handler for decks from Moxfield

type Handler struct{}

func (h Handler) Source() core.APISource {
	return core.SourceMoxfield
}

func (h Handler) Fetch(ctx context.Context, input string, cfg core.Config, ovrr core.DeckMeta) (core.Deck, error) {
	_ = ctx
	_ = input
	_ = cfg

	return core.Deck{}, nil
}

func (h Handler) Search(ctx context.Context, cfg core.Config, scfg core.SearchConfig) ([]core.DeckMeta, error) {
	_ = ctx
	_ = cfg

	return []core.DeckMeta{}, nil
}

func init() {
	core.RegisterSource(Handler{})
}
