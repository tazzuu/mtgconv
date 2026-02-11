package core

// base object class for holding global configs to be passed throughout the program
// this should generally be set via the cli interface
// and maps back to cli settings
type Config struct {
	Debug          bool   // enable debug entrypoint for dev testing
	Verbose        bool   // enable verbose logging
	PrintVersion   bool   // print version and quit
	CompatibilityMode bool // enable compatibility mode for modified output
	Version        string // the current version of the program TODO: replace this with BuildInfo
	InputSource InputSource
	OutputFilename string // output file name
	OutputDir string // output directory name
	OutputFormat   OutputFormat
	UserAgent      string // user agent token string for web requests
	UrlString      string // user supplied URL to query for decklist
	AutoFilename bool // automatically create an output filename
	SaveJSON bool // save a copy of the API response JSON
	Build BuildInfo // current build of the program
}

// create an empty Config with default values
// NOTE: This is intended for internal package usage; the CLI interface should update the config with different defaults!
func DefaultConfig(build BuildInfo) Config {
	return Config{
		// Version: false,
		InputSource: InputMoxfieldURL,
		OutputFilename: "auto",
		AutoFilename: true,
		OutputFormat: OutputDCK,
		Build: build,
	}
}

// config for doing deck searches
type SearchConfig struct {
	SortType SortType
	MinBracket CommanderBracket
	MaxBracket CommanderBracket
	SortDirection SortDirection
	DeckFormat DeckFormat
	Username string
	PageStart int
	PageEnd int
	PageSize int
}

// returns a Search Config with default settings
// NOTE: This is intended for internal package usage; the CLI interface should update the config with different defaults!
func DefaultSearchConfig() SearchConfig {
	config := SearchConfig{
		SortType: SortLikes,
		MinBracket: CommanderBracket1,
		MaxBracket: CommanderBracket5,
		SortDirection: SortDesc,
		DeckFormat: DeckFormatCommander,
		PageStart: 1,
		PageEnd: 1,
		PageSize: 64, // NOTE: some Search API's default to less or dont use this
	}
	return config
}