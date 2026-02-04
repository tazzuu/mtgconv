package core

import (
	"strconv"
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

// output data format
// NOTE: also used as file extension for auto output filename
type OutputFormat string
const (
	OutputDCK  OutputFormat = "dck"
	OutputJSON OutputFormat = "json"
	OutputTXT OutputFormat = "txt"
	// add more formats here
)
func OutputFormats() []OutputFormat {
	return []OutputFormat{OutputDCK, OutputJSON, OutputTXT}
}

// sections in a deck list
type BoardType string
const (
	BoardMain      BoardType = "main"
	BoardCommander BoardType = "commander"
	BoardSideboard BoardType = "sideboard"
	BoardMaybeboard BoardType = "maybeboard" // "Considering" cards
)

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