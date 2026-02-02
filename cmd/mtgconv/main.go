package main

// main cli entrypoint for the program
//
// USAGE:
// ./mtgconv --user-agent "$MOXKEY" convert --output-filename auto --save-json --compatibility-mode https://moxfield.com/decks/mrIbUS4YX0yqhpQMgsWY2w
//

import (
	// "flag"
	"fmt"
	// "log"
	"log/slog"
	// "os"

	"github.com/alecthomas/kong"

	"mtgconv/pkg/mtgconv2/core"
	_ "mtgconv/pkg/mtgconv2/all" // registers all the input/output handlers
)

// overwrite this at build time ;
// -ldflags="-X 'main.Version=someversion'"
// also gets set by goreleaser; https://goreleaser.com/cookbooks/using-main.version/
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)


// global cli options go here
type Context struct {
	Debug bool
	Verbose bool
	UserAgent string
	SaveJSON bool
}

type Convert struct {
	OutputFilename string `default:"-" help:"Output filename, use '-' to write to stdout, use 'auto' to automatically generate a filename based on decklist metadata"`
	OutputFormat string `default:"dck"`
	CompatibilityMode bool `default:"false"`
	Input string `arg:"" help:"file or URL path to input deck list"`
	}
func (c *Convert) Run(ctx *Context) error {
	core.ConfigureLogging(ctx.Verbose)

	// check the output format
	format, err := core.ParseOutputFormat(c.OutputFormat)
	if err != nil {
		// log.Fatalf("Error: invalid output format: %v", err)
		return &core.UnknownOutputFormat{Format: core.OutputFormat(c.OutputFormat)}
	}

	// create config object
	config := core.Config{
		Debug:          ctx.Debug,
		Verbose:        ctx.Verbose,
		OutputFilename: c.OutputFilename,
		UserAgent:      ctx.UserAgent,
		UrlString:      c.Input,
		OutputFormat: format,
		CompatibilityMode: c.CompatibilityMode,
		SaveJSON: ctx.SaveJSON,
	}

	if c.OutputFilename == "auto" {
		config.AutoFilename = true
	}

	// if we are doing debug run that instead and quit
	if ctx.Debug {
		slog.Debug("Starting DebugFunc from cmd/main.go")
		core.DebugFunc(config)
		return nil
	}

	// main entrypoint for the program
	err = core.RunCLI(config)
	if err != nil {
		// log.Fatalf("error running program: %v", err)
		return err
	}

	return nil
}

type Search struct {
	Input string `arg:"" default:"https://moxfield.com" help:"URL domain to search for decks"`
}
func (s *Search) Run(ctx *Context) error {
	core.ConfigureLogging(ctx.Verbose)
	slog.Debug("starting cli search")
	searchConfig, err := core.DefaultSearchConfig()
	if err != nil {
		return err
	}
	slog.Debug("got search config", "searchConfig", searchConfig)

	config := core.Config{
		UrlString: s.Input,
		Debug:          ctx.Debug,
		Verbose:        ctx.Verbose,
		UserAgent:      ctx.UserAgent,
		SaveJSON: ctx.SaveJSON,
	}
	slog.Debug("got config", "config", config)
	err = core.SearchCLI(config, searchConfig)
	if err != nil {
		return err
	}
	return nil
}

type Version struct {}
func (v *Version) Run(ctx *Context) error {
	fmt.Printf("%s", version)
	return nil
}

var cli struct {
	// global options
	Debug bool `help:"Enable debug mode."`
	UserAgent string `help:"user token to use for web requests" default:"default-user-agent"`
	Verbose bool `default:"false" help:"enable verbose logging"`
	SaveJSON bool `default:"false" help:"save API request JSON for inspection"`

	// subcommands
	Version Version `cmd:"" help:"Print version information and quit"`
	Convert Convert `cmd:"" help:"Convert a deck list to another format"`
	Search Search `cmd:"" help:"Search for decks to output"`
}

func main() {
  ctx := kong.Parse(&cli)
  // Call the Run() method of the selected parsed command.
  err := ctx.Run(&Context{
	UserAgent: cli.UserAgent,
	Debug: cli.Debug,
	Verbose: cli.Verbose,
	SaveJSON: cli.SaveJSON,
	})
  ctx.FatalIfErrorf(err)
}







// OLD CLI PARSING
// TODO: REMOVE THIS


// cli parser logic goes here
// USAGE:
// urlString, config := parse()
//
//	debug := config.Debug
//	verbose := config.Verbose
// func parseCLI() core.Config {
// 	verbose := flag.Bool("verbose", false, "enable verbose logging (log level DEBUG)")
// 	debug := flag.Bool("debug", false, "enable debug entrypoint")
// 	printVersion := flag.Bool("version", false, "print version and quit")
// 	outputFilename := flag.String("output", "-", "output filename (use - for stdout)")
// 	outputFormat := flag.String("output-fmt", "dck", "output format")
// 	userAgent := flag.String("user-agent", "foooo", "user token to use for web requests")
// 	compatibilityMode := flag.Bool("compat", false, "enable compatibility mode for output")
// 	saveJSON := flag.Bool("save-json", false, fmt.Sprintf("save a copy of the response json to %s", core.ResponseJSONFilename))
// 	flag.Parse()

// 	posArgs := flag.Args() // all positional args passed

// 	// check if version was passed
// 	// NOTE: had to put the check here to avoid requiring pos arg to -version flag
// 	if *printVersion {
// 		PrintVersionAndQuit()
// 	}

// 	var urlString string

// 	if len(posArgs) < 1 {
// 		log.Fatalf("Error: requires at least 1 positional argument for urlString")
// 	} else {
// 		urlString = posArgs[0]
// 	}

// 	// check the output format
// 	format, err := core.ParseOutputFormat(*outputFormat)
// 	if err != nil {
// 		log.Fatalf("Error: invalid output format: %v", err)
// 	}

// 	// create config object
// 	config := core.Config{
// 		Debug:          *debug,
// 		Verbose:        *verbose,
// 		PrintVersion:   *printVersion,
// 		OutputFilename: *outputFilename,
// 		Version:        version,
// 		UserAgent:      *userAgent,
// 		UrlString:      urlString,
// 		OutputFormat: format,
// 		CompatibilityMode: *compatibilityMode,
// 		SaveJSON: *saveJSON,
// 	}

// 	if *outputFilename == "auto" {
// 		config.AutoFilename = true
// 	}

// 	return config
// }

// func PrintVersionAndQuit() {
// 	fmt.Printf("%s", version)
// 	os.Exit(0)
// }

// NOTE: log.Fatalf only in the cli interface ; 'return fmt.Errorf' in all other interfaces
// func main2() {
// 	// get the cli args
// 	config := parseCLI()
// 	debug := config.Debug
// 	verbose := config.Verbose

// 	// start logging
// 	core.ConfigureLogging(verbose)

// 	// if we are doing debug run that instead and quit
// 	if debug {
// 		slog.Debug("Starting DebugFunc from cmd/main.go")
// 		core.DebugFunc(config)
// 		return
// 	}

// 	// main entrypoint for the program
// 	err := core.RunCLI(config)
// 	if err != nil {
// 		log.Fatalf("error running program: %v", err)
// 	}
// }
