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

	err = CreateOutput(
		output,
		config.OutputFilename,
		config.OutputDir,
		config.AutoFilename,
		deck.Meta.Name,
		deck.Meta.Version,
		config.OutputFormat.GetExtension(),
	)

	return err
}

// decide where to print the output
func CreateOutput(contents string, outputFilename string, outputDir string, autoFilename bool, deckName string, deckVersion int, fileExtension string) error {
	var out *os.File
	var err error
	// print to stdout if - or empty string passed
	if outputFilename == "-" || outputFilename == "" {
		out = os.Stdout
	} else {
		// check if output directory was passed
		if outputDir != "" {
			// make the output dir
			slog.Debug("creating output directory", "OutputDir", outputDir)
			err := os.MkdirAll(outputDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		// auto generate a filename
		var cleanedOutputFilename string = outputFilename
		if autoFilename == true {
			cleanedOutputFilename = GenerateSafeFilename(deckName, deckVersion, fileExtension)
		}
		// add the output dir name
		if outputDir != "" {
			cleanedOutputFilename = filepath.Join(outputDir, cleanedOutputFilename)
		}
		slog.Debug("resolved final output filename", "outputFilename", cleanedOutputFilename)
		out, err = os.Create(cleanedOutputFilename)
		if err != nil {
			return err
		}
		// NOTE: this attempt to return error from defer is tricky and suspicious
		defer func() {
			if cerr := out.Close(); cerr != nil {
				err = cerr
			}
		}()
	}
	slog.Info("saving to output file", "out", out.Name())
	if _, err := fmt.Fprintln(out, contents); err != nil {
		return err
	}

	return nil
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
