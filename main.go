package main

import (
	"fmt"
	"os"
	"encoding/json"
)

// USAGE:
// go run . --user-agent "$MOXKEY" https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ
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
	// debug := config.Debug
	verbose := config.Verbose

	configureLogging(verbose)



	// if we are doing debug run that instead and quit
	// if debug {
	// 	DebugFunc(urlString, config)
	// 	return
	// }

	// make sure we can connect to external resources and API's
	// err := CheckConnectivity()
	// if err != nil {
	// 	log.Fatalf("checking API connectivity: %v", err)
	// }

	deckID := deckIDFromURL(config.UrlString)
	deckAPIUrl := makeAPIUrl(deckID)

	jsonStr, err := fetchJSON(deckAPIUrl, config.UserAgent)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// fmt.Println(jsonStr)

	var deck DeckResponse
	if err := json.Unmarshal([]byte(jsonStr), &deck); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse JSON response: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(deck.Authors)

}