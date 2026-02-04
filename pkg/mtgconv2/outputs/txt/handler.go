package txt

import (
	_ "embed"
	"fmt"
	"log/slog"
	"mtgconv/pkg/mtgconv2/core"
	"strings"
	"text/template"
)

//go:embed templates/output.txt
var dckTemplateStr string

type Handler struct{}

func (h Handler) Format() core.OutputFormat {
	return core.OutputTXT
}

func (h Handler) Render(deck core.Deck, cfg core.Config) (string, error) {
	_ = deck
	_ = cfg
	slog.Debug("running TXT output handler")

	funcMap := template.FuncMap{
		"CollectCommanders": core.CollectCommanders,
		"CollectMainboard": core.CollectMainboard,
		"FormatTxtLine": func(entry core.DeckEntry) string {
			return fmt.Sprintf("%d %v", entry.Quantity, entry.Card.Name)
		},
	}

	tmpl, err := template.New("txt").Funcs(funcMap).Parse(dckTemplateStr)
	if err != nil {
		slog.Error("Error initializing report template", "err", err)
		return "", &core.TemplateInitializationError{Message: err}
	}

	var output strings.Builder
	if err := tmpl.Execute(&output, deck); err != nil {
		slog.Error("Error creating report template", "err", err)
		return "", &core.TemplateExecutionError{Message: err}
	}

	result := output.String()

	return result, nil
}

func init() {
	core.RegisterOutput(Handler{})
}
