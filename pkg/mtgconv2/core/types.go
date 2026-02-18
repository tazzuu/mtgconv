package core

import (
	"strconv"
	"strings"
)

// object type to use for passing the program build info to internal components
type BuildInfo struct{ Version, Commit, Date, Program string }

// designation for various API domains that we can query
type APISource string
const (
	SourceMoxfield APISource = "moxfield.com"
	SourceArchidekt APISource = "archidekt.com"
	// add more sources here
)
func APISources() []APISource {
	return []APISource{SourceMoxfield, SourceArchidekt}
}
func ParseAPISource(hostname string) (APISource, InputSource,  error) {
	switch APISource(hostname) {
	case SourceMoxfield:
		return SourceMoxfield, InputMoxfieldURL, nil
	case SourceArchidekt:
		return SourceArchidekt, InputArchidektURL, nil
	// add more cases here
	default:
		return "", "", &UnrecognizedDomain{Message: hostname}
	}
}

type InputSourceType string
const (
	InputSourceTypeFile InputSourceType = "file"
	InputSourceTypeURL InputSourceType = "url"
)

// all of the supported input data formats
type InputSource string
const (
	InputMoxfieldURL InputSource = "moxfield-url"
	InputArchidektURL InputSource = "archidekt-url"
	InputShinyCSV InputSource = "shiny-csv"
	InputTxtMoxfield InputSource = "txt-moxfield"
)
// return the InputSourceType and if its a file type or not
func (f InputSource) Type() (InputSourceType, bool) {
	switch f {
	// url types
	case InputMoxfieldURL, InputArchidektURL:
		return InputSourceTypeURL, false
	// file types
	case InputShinyCSV, InputTxtMoxfield:
		return InputSourceTypeFile, true
	default:
		return "", false
	}
}
func InputSources() []InputSource {
	return []InputSource{
		InputMoxfieldURL,
		InputArchidektURL,
		InputShinyCSV,
		InputTxtMoxfield,
	}
}
func ParseInputSource(raw string) (InputSource, error) {
	switch InputSource(raw) {
	case InputMoxfieldURL:
		return InputMoxfieldURL, nil
	case InputArchidektURL:
		return InputArchidektURL, nil
	case InputShinyCSV:
		return InputShinyCSV, nil
	case InputTxtMoxfield:
		return InputTxtMoxfield, nil
	default:
		return "", &UnknownInputSource{Source: InputSource(raw)}
	}
}



// output data format
type OutputFormat string
const (
	OutputDCK  OutputFormat = "dck"
	OutputJSON OutputFormat = "json"
	OutputTXT OutputFormat = "txt"
	OutputMoxfieldCollection OutputFormat = "moxfield-collection"
	// add more formats here
)
// return the file extension to use
func (f OutputFormat) GetExtension() string {
	label := strings.ToLower(string(f))
	switch label {
	case string(OutputDCK):
		return ".dck"
	case string(OutputTXT):
		return ".txt"
	case string(OutputJSON):
		return ".json"
	case string(OutputMoxfieldCollection):
		return ".moxfield-collection.csv"
	default:
		return "." + string(f)
	}
}
// determine the correct output format based on the supplied string
func ParseOutputFormat(raw string) (OutputFormat, error) {
	switch strings.ToLower(raw) {
	case string(OutputDCK):
		return OutputDCK, nil
	case string(OutputTXT):
		return OutputTXT, nil
	case string(OutputJSON):
		return OutputJSON, nil
	case string(OutputMoxfieldCollection):
		return OutputMoxfieldCollection, nil
	default:
		return "", &UnknownOutputFormat{OutputFormat(raw)}
	}
}
// return all registered output formats
func OutputFormats() []OutputFormat {
	return []OutputFormat{OutputDCK, OutputJSON, OutputTXT,OutputMoxfieldCollection}
}

// sections in a deck list
type BoardType string
const (
	BoardMain      BoardType = "main"
	BoardCommander BoardType = "commander"
	BoardSideboard BoardType = "sideboard"
	BoardMaybeboard BoardType = "maybeboard" // "Considering" cards
)
func BoardTypes() []BoardType {
	return []BoardType{BoardMain, BoardCommander, BoardSideboard, BoardMaybeboard}
}

// card finishes
type FinishType string
const (
	FinishDefault FinishType = "default"
	FinishFoil    FinishType = "foil"
	FinishNonfoil FinishType = "nonfoil"
	FinishEtched  FinishType = "etched"
)

// query sorting methods
type SortType string
const (
	SortLikes SortType = "likes"
	SortViews SortType = "views"
	SortUpdated SortType = "updated"
)
func SortTypes() []SortType {
	return []SortType{
		SortLikes, SortViews, SortUpdated,
	}
}

type CommanderBracket int
func (c CommanderBracket) String() string {
	return strconv.Itoa(int(c))
}
const (
	CommanderBracket1 CommanderBracket = 1
	CommanderBracket2 CommanderBracket = 2
	CommanderBracket3 CommanderBracket = 3
	CommanderBracket4 CommanderBracket = 4
	CommanderBracket5 CommanderBracket = 5
)
func CommanderBrackets() []CommanderBracket {
	return []CommanderBracket{
		CommanderBracket1, CommanderBracket2, CommanderBracket3, CommanderBracket4, CommanderBracket5,
	}
}


type SortDirection string
const (
	SortAsc SortDirection = "ascending"
	SortDesc SortDirection = "descending"
)
func SortDirections() []SortDirection {
	return []SortDirection{SortAsc, SortDesc}
}

type DeckFormat string
const (
	DeckFormatCommander = "commander"
	DeckFormatCommanderPrecons = "commanderPrecons" // NOTE: Moxfield only
)
func DeckFormats() []DeckFormat {
	return []DeckFormat{DeckFormatCommander, DeckFormatCommanderPrecons}
}