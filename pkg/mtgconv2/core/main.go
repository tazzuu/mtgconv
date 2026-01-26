package core

import (
	"fmt"
	"os"
	"log/slog"
	"context"
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
		// auto generate a filename
		var outputFilename string = config.OutputFilename
		if config.AutoFilename == true {
			outputFilename = GenerateSafeFilename(config, deck)
		}
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
	if _, err := fmt.Fprintln(out, output); err != nil {
		return err
	}

	return err
}