package mtgconv

import (
	"fmt"
	"log/slog"
)

func DebugFunc(config Config) {
	slog.Debug("Running DebugFunc")
	slog.Debug("Got config", "config", config)

	// parse provided config
	slog.Debug("Starting Deck loader")
	LoadDeck(config)

}

func DebugFunc2(config Config) {
	// put methods here for debug and development
	slog.Debug("Running DebugFunc2")
	slog.Debug("Got config", "config", config)

	// NOTE: assuming that the input here is a URL to a website

	// figure out which site the input is hosted on
	apiSource, err := DetectURLSource(config.UrlString)
	if err != nil {
		slog.Error("error detecting domain from URL", "url", config.UrlString)
		// TODO: How to halt the program here and quit with error?
	}

	slog.Debug("Got URL API source", "apiSource", apiSource)
	config.SetDomain(apiSource)
	slog.Debug("Updated config", "config", config)

	slog.Debug("getting API response for domain")


	// get the API JSON response
	// jsonStr, err := FetchMoxfieldDecklistJson(config)
	// if err != nil {
	// 	slog.Error("error getting deck ID from URL: %v", err)
	// }

	// // convert it to internal Moxfield Deck object type
	// moxDeck, err := MakeMoxfieldDeck(jsonStr)
	// if err != nil {
	// 	slog.Error("error parsing the Moxfield API response JSON", "err", err)
	// }

	// convert to final .dck decklist format
	// result, err := MoxfieldDeckToDckFormat(moxDeck)
	// if err != nil {
	// 	slog.Error("error while generating the .dck template", "err", err)
	// }
	// fmt.Println(result)

	fmt.Println("")
}
