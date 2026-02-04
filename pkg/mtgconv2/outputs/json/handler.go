package json

import (
	"log/slog"
	"encoding/json"
	"mtgconv/pkg/mtgconv2/core"
)

type Handler struct{}

func (h Handler) Format() core.OutputFormat {
	return core.OutputJSON
}

func (h Handler) Render(deck core.Deck, cfg core.Config) (string, error) {
	_ = deck
	_ = cfg
	slog.Debug("running JSON output handler")

	// Serialize struct to JSON.
	deckJSON, err := json.Marshal(deck)
	if err != nil {
		return "", err
	}

	betterJSON, err := core.PrettyJSON(string(deckJSON))

	return betterJSON, nil
}

func init() {
	core.RegisterOutput(Handler{})
}
