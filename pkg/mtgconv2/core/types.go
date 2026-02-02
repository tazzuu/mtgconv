package core

import (
	"strconv"
)

// designation for various API domains that we can query
type APISource string

const (
	SourceMoxfield APISource = "moxfield.com"
	// add more sources here
)

// output data format
// NOTE: also used as file extension for auto output filename
type OutputFormat string

const (
	OutputDCK  OutputFormat = "dck"
	OutputJSON OutputFormat = "json"
	// add more formats here
)

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

type CommanderBracket int

const (
	CommanderBracket1 CommanderBracket = 1
	CommanderBracket2 CommanderBracket = 2
	CommanderBracket3 CommanderBracket = 3
	CommanderBracket4 CommanderBracket = 4
	CommanderBracket5 CommanderBracket = 5
)

func (c CommanderBracket) String() string {
	return strconv.Itoa(int(c))
}

type SortDirection string

const (
	SortAsc SortDirection = "ascending"
	SortDesc SortDirection = "descending"
)

type DeckFormat string

const (
	DeckFormatCommander = "commander"
)