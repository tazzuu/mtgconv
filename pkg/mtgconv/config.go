package mtgconv

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




