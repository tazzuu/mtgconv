package main

// main cli entrypoint for the program
// USAGE:
// go run cmd/mtgconv/*.go --user-agent "$MOXKEY" https://moxfield.com/decks/Wrcumkgcc0qjIB2bwoDvqQ
// https://api.moxfield.com/v2/decks/all/1RAZHbA3H0WE7EHFK36-_Q
//

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"mtgconv/pkg/mtgconv"
)

// overwrite this at build time ;
// -ldflags="-X 'main.Version=someversion'"
// also gets set by goreleaser; https://goreleaser.com/cookbooks/using-main.version/
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// cli parser logic goes here
// USAGE:
// urlString, config := parse()
//
//	debug := config.Debug
//	verbose := config.Verbose
func parseCLI() mtgconv.Config {
	verbose := flag.Bool("verbose", false, "enable verbose logging (log level DEBUG)")
	debug := flag.Bool("debug", false, "enable debug entrypoint")
	printVersion := flag.Bool("version", false, "print version and quit")
	outputFilename := flag.String("output", "-", "output filename (use - for stdout)")
	userAgent := flag.String("user-agent", "foooo", "user token to use for web requests")
	flag.Parse()

	posArgs := flag.Args() // all positional args passed

	// check if version was passed
	// NOTE: had to put the check here to avoid requiring pos arg to -version flag
	if *printVersion {
		PrintVersionAndQuit()
	}

	var urlString string

	if len(posArgs) < 1 {
		log.Fatalf("Error: requires at least 1 positional argument for urlString")
	} else {
		urlString = posArgs[0]
	}

	// create config object
	config := mtgconv.Config{
		Debug:          *debug,
		Verbose:        *verbose,
		PrintVersion:   *printVersion,
		OutputFilename: *outputFilename,
		Version:        version,
		UserAgent:      *userAgent,
		UrlString:      urlString,
	}

	return config
}

func PrintVersionAndQuit() {
	fmt.Printf("%s", version)
	os.Exit(0)
}

// NOTE: log.Fatalf only in the cli interface ; 'return fmt.Errorf' in all other interfaces
func main() {
	// get the cli args
	config := parseCLI()
	debug := config.Debug
	verbose := config.Verbose
	outputFilename := config.OutputFilename
	mtgconv.ConfigureLogging(verbose)

	// if we are doing debug run that instead and quit
	if debug {
		slog.Debug("Running DebugFunc")
		mtgconv.DebugFunc()
		return
	}

	// make sure we can connect to external resources and API's
	// TODO: implement this
	slog.Debug("Checking API Conectivity")
	err := mtgconv.CheckConnectivity()
	if err != nil {
		log.Fatalf("checking API connectivity: %v", err)
	}

	deck, err := mtgconv.MoxfieldURLtoDckFormat(config)
	if err != nil {
		log.Fatalf("error converting to .dck declist format: %v", err)
	}

	// write to stdout or to file
	var out *os.File
	if outputFilename == "-" {
		out = os.Stdout
	} else {
		out, err = os.Create(outputFilename)
		if err != nil {
			log.Fatalf("creating output file: %v", err)
		}
		defer func() {
			if err := out.Close(); err != nil {
				log.Printf("closing output file: %v", err)
			}
		}()
	}

	if _, err := fmt.Fprintln(out, deck); err != nil {
		log.Fatalf("writing output: %v", err)
	}
}
