package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"mtgconv/pkg/mtgconv"
)

// cli parsing logic for the program

// cli parser logic goes here
// USAGE:
// urlString, config := parse()
// 	debug := config.Debug
// 	verbose := config.Verbose
func parseCLI() (mtgconv.Config) {
	verbose := flag.Bool("verbose", false, "enable verbose logging (log level DEBUG)")
	debug := flag.Bool("debug", false, "enable debug entrypoint")
	printVersion := flag.Bool("version", false, "print version and quit")
	outputFilename := flag.String("output", "output", "filename prefix for report output")
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
		UserAgent: *userAgent,
		UrlString: urlString,
	}

	return config
}

func PrintVersionAndQuit() {
	fmt.Printf("%s", version)
	os.Exit(0)
}


