package core

import (
	"fmt"
	"os"
	"log/slog"
	"context"
	"path/filepath"
)

// main entrypoint for the program when running from the cli
// TODO: move this back into the cmd/main.go instead
// TODO: API connectivity check
func RunCLI(config Config, deckMetaOverride DeckMeta) (err error) {
	slog.Debug("got config", "config", config)

	// run the main pipeline with the given config
	slog.Info("Starting deck import pipeline", "input", config.UrlString)
	output, deck, err := Run(context.Background(), config, deckMetaOverride)
	if err != nil {
		slog.Error("error running deck processing pipeline", "err", err)
		return err
	}

	// "-" or "" means write the deck to stdout; anything else is a real file path
	toStdout := config.OutputFilename == "-" || config.OutputFilename == ""

	// resolve the final on-disk path for the deck output
	deckPath := ResolveOutputPath(
		config.OutputFilename,
		config.OutputDir,
		config.AutoFilename,
		deck.Meta.Name,
		deck.Meta.Version,
		deck.Meta.Bracket,
		config.OutputFormat.GetExtension(),
	)

	// write the converted deck (stdout or file)
	if err := CreateOutput(output, deckPath, toStdout); err != nil {
		return err
	}

	// write the primer as a ".primer.md" sidecar next to the deck file.
	// skipped when: --primer skip, writing the deck to stdout, or no primer exists.
	if config.Primer != PrimerSkip && !toStdout && deck.Primer.Content != "" {
		primerPath := PrimerSidecarPath(deckPath)
		slog.Info("saving deck primer", "path", primerPath)
		if err := SaveTxtToFile(primerPath, deck.Primer.Content); err != nil {
			return err
		}
	}

	return nil
}

// ResolveOutputPath computes the final on-disk path for a deck output file.
// It does NOT handle the stdout case ("-"): callers decide that separately.
// When autoFilename is set the name is generated from the deck metadata;
// otherwise outputFilename is used verbatim. outputDir, if set, is prepended.
func ResolveOutputPath(outputFilename, outputDir string, autoFilename bool, deckName string, deckVersion int, deckBracket CommanderBracket, ext string) string {
	name := outputFilename
	if autoFilename {
		name = GenerateSafeFilename(deckName, deckVersion, deckBracket, ext)
	}
	if outputDir != "" {
		name = filepath.Join(outputDir, name)
	}
	return name
}

// PrimerSidecarPath derives the primer markdown path that pairs with a deck
// output path by swapping the extension for ".primer.md".
// e.g. "decks/Yshtola.b3.v5.dck" -> "decks/Yshtola.b3.v5.primer.md"
func PrimerSidecarPath(deckPath string) string {
	ext := filepath.Ext(deckPath)
	return deckPath[:len(deckPath)-len(ext)] + ".primer.md"
}

// CreateOutput writes the rendered deck to stdout or to a resolved file path,
// creating the parent directory as needed.
func CreateOutput(contents string, resolvedPath string, toStdout bool) error {
	if toStdout {
		_, err := fmt.Fprintln(os.Stdout, contents)
		return err
	}
	// ensure the parent directory exists
	if dir := filepath.Dir(resolvedPath); dir != "" && dir != "." {
		slog.Debug("creating output directory", "dir", dir)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	slog.Info("saving to output file", "path", resolvedPath)
	return SaveTxtToFile(resolvedPath, contents)
}

// TODO: move this somewhere else
func SearchCLI(config Config, searchConfig SearchConfig) error {
	ctx := context.Background()
	slog.Debug("starting processing pipeline")

	slog.Debug("configuring source handler", "config.InputSource", config.InputSource)
	sourceHandler, err := HandlerForSource(config.InputSource)
	if err != nil {
		return err
	}

	slog.Debug("configuring search settings")

	slog.Info("searching for decks", "source", config.UrlString)
	result, err := sourceHandler.Search(ctx, config, searchConfig)
	if err != nil {
		return err
	}
	slog.Info("got search results", "n", len(result))

	slog.Debug("retrieving deck list for each search result")
	for i, entry := range result {
		// update the main config with some default values
		// TODO: find better way to implement this
		newConfig := config
		newConfig.AutoFilename = true
		newConfig.OutputFilename = "auto"
		newConfig.CompatibilityMode = true
		newConfig.OutputFormat = OutputDCK
		newConfig.UrlString = entry.URL
		slog.Debug("retrieving deck", "i", i, "name", entry.Name, "url", entry.URL)
		// NOTE: need to inject some extra meta because some meta is only returned by Search and not Fetch
		err := RunCLI(newConfig, DeckMeta{Bracket: entry.Bracket})
		if err != nil {
			return err
		}
	}

	return nil
}
