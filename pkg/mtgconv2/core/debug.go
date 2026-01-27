package core

import (
	"fmt"
	"log/slog"
	"context"
)

// use this space to put methods I am developing and testing to trigger from the cli
func DebugFunc(config Config) {
	slog.Debug("Running DebugFunc mtgconv2.core", "config", config)
	ctx := context.Background()
	slog.Debug("starting processing pipeline")
	src, err := DetectURLSource(config.UrlString)
	if err != nil {
		// return "", Deck{}, err
		return
	}

	slog.Debug("configuring source handler")
	sourceHandler, err := HandlerForSource(src)
	if err != nil {
		// return "", Deck{}, err
		return
	}

	slog.Debug("configuring search settings")
	searchConfig, err := DefaultSearchConfig()
	slog.Debug("fetching data from source")
	data, err := sourceHandler.Search(ctx, config, searchConfig)
	fmt.Println(data)





	// _, err := sourceHandler.Fetch(ctx, config.UrlString, config)
	// if err != nil {
	// 	// return "", Deck{}, err
	// 	return
	// }
	// output, err := Run(context.Background(), config)
	// if err != nil {
	// 	slog.Error("error running deck processing pipeline", "err", err)
	// 	return
	// }
	// fmt.Println(output)
}