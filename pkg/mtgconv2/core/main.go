package core

import (
	"fmt"
	"log/slog"
	"context"
)

// main entrypoint for the program when running from the cli
func RunCLI(config Config) error {
	slog.Debug("Running RunCLI mtgconv2.core", "config", config)
	output, err := Run(context.Background(), config)
	if err != nil {
		slog.Error("error running deck processing pipeline", "err", err)
		return err
	}
	fmt.Println(output)

	return nil
}