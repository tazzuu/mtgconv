package dck

import (
	"strings"
	"log/slog"

	_ "embed"
	"text/template"

	"mtgconv/pkg/mtgconv2/core"
)

// handler for the .dck decklist output format

//go:embed templates/dck.txt
var dckTemplateStr string

type Handler struct{}

func (h Handler) Format() core.OutputFormat {
	return core.OutputDCK
}

func (h Handler) Render(deck core.Deck, cfg core.Config) (string, error) {
	_ = deck
	_ = cfg

	var result string
	var compatibilityMode bool = cfg.CompatibilityMode

	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ConvertDate": core.ConvertDate,
		"CollectAuthors": core.CollectAuthors,
		"CollectCommanders": core.CollectCommanders,
		"CollectMainboard": core.CollectMainboard,
		"CollectSideboard": func(d core.Deck) []core.DeckEntry {
			entries := core.CollectSideboard(d)
			if compatibilityMode && len(entries) > 10 {
				return entries[:10]
			}
			return entries
		},
		"FormatDckLine": func(entry core.DeckEntry) string {
			return FormatDckLine(entry, compatibilityMode)
		},
		// cheat in some config meta values
		"ProgramName": func() string {
			return cfg.Build.Program
		},
		"ProgramVersion": func() string {
			return cfg.Build.Version
		},
	}

	tmpl, err := template.New("dck").Funcs(funcMap).Parse(dckTemplateStr)
	if err != nil {
		slog.Error("Error initializing report template", "err", err)
		return "", &core.TemplateInitializationError{Message: err}
	}

	var output strings.Builder
	if err := tmpl.Execute(&output, deck); err != nil {
		slog.Error("Error creating report template", "err", err)
		return "", &core.TemplateExecutionError{Message: err}
	}

	result = output.String()

	return result, nil
	// return "", fmt.Errorf("dck output handler not implemented")
}

func init() {
	core.RegisterOutput(Handler{})
}
