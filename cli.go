package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// cli parsing logic for the program

// base object class for holding global configs to be passed throughout the program
// this should generally be set via the cli interface
// and maps back to cli settings
type Config struct {
	Debug          bool   // enable debug entrypoint for dev testing
	Verbose        bool   // enable verbose logging
	PrintVersion   bool   // print version and quit
	Version        string // the current version of the program
	OutputFilename string // output file name
	UserAgent string // user agent token string for web requests
	UrlString string // user supplied URL to query for decklist
}

// cli parser logic goes here
// USAGE:
// urlString, config := parse()
// 	debug := config.Debug
// 	verbose := config.Verbose
func parseCLI() (Config) {
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
	config := Config{
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


