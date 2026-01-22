package dck

import (
	"fmt"

	"mtgconv/pkg/mtgconv2/core"
)

type Handler struct{}

func (h Handler) Format() core.OutputFormat {
	return core.OutputDCK
}

func (h Handler) Render(deck core.Deck, cfg core.Config) (string, error) {
	_ = deck
	_ = cfg
	return "", fmt.Errorf("dck output handler not implemented")
}

func init() {
	core.RegisterOutput(Handler{})
}
