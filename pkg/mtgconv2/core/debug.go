package core

import (
	"fmt"
	"log/slog"
	"context"
)

func DebugFunc(config Config) {
	slog.Debug("Running DebugFunc mtgconv2.core", "config", config)
	output, err := Run(context.Background(), config)
	if err != nil {
		slog.Error("error running deck processing pipeline", "err", err)
		return
	}
	fmt.Println(output)
}