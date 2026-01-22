package core

import (
	"fmt"
	"os"
	"log/slog"
	"context"
)

// main entrypoint for the program when running from the cli
func RunCLI(config Config) (err error) {
	slog.Debug("Running RunCLI mtgconv2.core", "config", config)
	// run the main pipeline with the given config
	output, err := Run(context.Background(), config)
	if err != nil {
		slog.Error("error running deck processing pipeline", "err", err)
		return err
	}

	// decide where to print the output
	var out *os.File
	if config.OutputFilename == "-" || config.OutputFilename == "" {
		out = os.Stdout
	} else {
		out, err = os.Create(config.OutputFilename)
		if err != nil {
			return err
		}
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