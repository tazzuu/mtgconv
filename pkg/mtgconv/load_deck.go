package mtgconv

import (
	"log/slog"
)

func LoadDeck(config Config)  { //DeckReader
	// check if a URL was passed as input
	var emptyStr string
	if config.UrlString != emptyStr {
		slog.Debug("detected passed URL", "url", config.UrlString)
		slog.Debug("identifying URL resource type")
		DetectURLSource(config.UrlString)
	}

	// TODO: load deck from file or stdin if that was passed instead
}