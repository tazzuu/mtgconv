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
func RunCLI(config Config) (err error) {
	slog.Debug("Running RunCLI mtgconv2.core", "config", config)
	// run the main pipeline with the given config
	output, deck, err := Run(context.Background(), config)
	if err != nil {
		slog.Error("error running deck processing pipeline", "err", err)
		return err
	}

	// decide where to print the output
	var out *os.File
	// print to stdout if - or empty string passed
	if config.OutputFilename == "-" || config.OutputFilename == "" {
		out = os.Stdout
	} else {
		// check if output directory was passed
		if config.OutputDir != "" {
			// make the output dir
			slog.Debug("creating output directory", "OutputDir", config.OutputDir)
			err := os.MkdirAll(config.OutputDir, os.ModePerm)
			if err != nil {
				return err
			}
		}
		// auto generate a filename
		var outputFilename string = config.OutputFilename
		if config.AutoFilename == true {
			outputFilename = GenerateSafeFilename(config, deck)
		}
		// add the output dir name
		if config.OutputDir != "" {
			outputFilename = filepath.Join(config.OutputDir, outputFilename)
		}
		slog.Debug("resolved final output filename", "outputFilename", outputFilename)
		out, err = os.Create(outputFilename)
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
	slog.Debug("saving to output file", "out", out.Name())
	if _, err := fmt.Fprintln(out, output); err != nil {
		return err
	}

	return err
}

// TODO: move this somewhere else
func SearchCLI(config Config, searchConfig SearchConfig) error {
	ctx := context.Background()
	slog.Debug("starting processing pipeline")
	src, err := DetectURLSource(config.UrlString)
	if err != nil {
		return err
	}

	slog.Debug("configuring source handler")
	sourceHandler, err := HandlerForSource(src)
	if err != nil {
		return err
	}

	slog.Debug("configuring search settings")

	slog.Debug("fetching data from source")
	result, err := sourceHandler.Search(ctx, config, searchConfig)
	if err != nil {
		return err
	}

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
		err := RunCLI(newConfig)
		if err != nil {
			return err
		}
	}

	return nil
}