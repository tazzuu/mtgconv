package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"mtgconv/pkg/mtgconv"
)

// USAGE:
// go run cmd/mtgconv/*.go --user-agent "$MOXKEY" https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ
// https://api.moxfield.com/v2/decks/all/1RAZHbA3H0WE7EHFK36-_Q
//


// overwrite this at build time ;
// -ldflags="-X 'main.Version=someversion'"
// also gets set by goreleaser; https://goreleaser.com/cookbooks/using-main.version/
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// NOTE: log.Fatalf only in the cli interface ; 'return fmt.Errorf' in all other interfaces
func main() {
	// get the cli args
	config := parseCLI()
	debug := config.Debug
	verbose := config.Verbose
	mtgconv.ConfigureLogging(verbose)



	// if we are doing debug run that instead and quit
	if debug {
		slog.Debug("Running DebugFunc")
		mtgconv.DebugFunc()
		return
	}

	// make sure we can connect to external resources and API's
	slog.Debug("Checking API Conectivity")
	err := mtgconv.CheckConnectivity()
	if err != nil {
		log.Fatalf("checking API connectivity: %v", err)
	}

	// get the deck ID from the provided URL
	slog.Debug("getting the deck ID from the provided url", "url", config.UrlString)
	deckID := mtgconv.DeckIDFromURL(config.UrlString)
	// create the API query URL
	deckAPIUrl := mtgconv.MakeAPIUrl(deckID)
	// fetch the JSON query result
	jsonStr, err := mtgconv.FetchJSON(deckAPIUrl, config.UserAgent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// fmt.Println(jsonStr)

	// convert the JSON string into Go objects
	deck := mtgconv.MakeMoxfieldDeckResponse(jsonStr)
	fmt.Println(deck.Authors)

}