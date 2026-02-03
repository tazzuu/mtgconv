package main

// main cli entrypoint for the program
//
// USAGE:
// ./mtgconv --user-agent "$MOXKEY" convert --output-filename auto --save-json --compatibility-mode https://moxfield.com/decks/mrIbUS4YX0yqhpQMgsWY2w
//

import (
	"fmt"
	"log/slog"

	"github.com/alecthomas/kong"

	"mtgconv/pkg/mtgconv2/core"
	_ "mtgconv/pkg/mtgconv2/all" // registers all the input/output handlers
)

// global cli options go here ; add them to the 'cli' struct below
// NOTE: dont add default values here, add them to 'cli'
type Context struct {
	Debug bool
	Verbose bool
	UserAgent string
	SaveJSON bool
	OutputDir string
	CompatibilityMode bool
}

// subcommand for converting a single deck into a different format
type Convert struct {
	OutputFilename string `default:"-" help:"Output filename, use '-' to write to stdout, use 'auto' to automatically generate a filename based on decklist metadata"`
	OutputFormat string `default:"dck"`
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
	config := ApplyConfig(*ctx)
	config.OutputFilename = c.OutputFilename
	config.UrlString = c.Input
	config.OutputFormat = format

	if config.OutputFilename == "auto" {
		config.AutoFilename = true
	}

	// if we are doing debug run that instead and quit
	if ctx.Debug {
		slog.Debug("Starting DebugFunc from cmd/main.go")
		core.DebugFunc(config)
		return nil
	}

	// main entrypoint for the program
	err = core.RunCLI(config, core.DeckMeta{})
	if err != nil {
		// log.Fatalf("error running program: %v", err)
		return err
	}

	return nil
}

// subcommand for searching for decks and outputting them in a specified format
type Search struct {
	Input string `arg:"" default:"https://moxfield.com" help:"URL domain to search for decks"`
}
func (s *Search) Run(ctx *Context) error {
	core.ConfigureLogging(ctx.Verbose)
	slog.Debug("starting cli search")
	searchConfig := core.DefaultSearchConfig()
	slog.Debug("got search config", "searchConfig", searchConfig)

	// create config object
	config := ApplyConfig(*ctx)
	config.UrlString = s.Input

	slog.Debug("got config", "config", config)
	err := core.SearchCLI(config, searchConfig)
	if err != nil {
		return err
	}
	return nil
}

// subcommand to print the program version information
type Version struct {}
func (v *Version) Run(ctx *Context) error {
	fmt.Printf("%s", BuildInfo.Version)
	return nil
}

// implements attributes that will be used in Context and the cli interface
var cli struct {
	// global options
	Debug bool `help:"Enable debug mode."`
	UserAgent string `help:"user token to use for web requests" default:"default-user-agent"`
	Verbose bool `default:"false" help:"enable verbose logging"`
	SaveJSON bool `default:"false" help:"save API request JSON for inspection"`
	OutputDir string `default:"converted-decks" help:"output directory name"`
	CompatibilityMode bool `default:"false" help:"apply compatibility formatting for deck list output formats where applicable to help when importing the deck lists into various programs"`

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
	OutputDir: cli.OutputDir,
	CompatibilityMode: cli.CompatibilityMode,
	})
  ctx.FatalIfErrorf(err)
}

