package core

type APISource string

const (
	SourceMoxfield APISource = "moxfield.com"
	// add more sources here
)

type OutputFormat string

const (
	OutputDCK  OutputFormat = "dck"
	OutputJSON OutputFormat = "json"
	// add more formats here
)

type BoardType string

const (
	BoardMain      BoardType = "main"
	BoardCommander BoardType = "commander"
	BoardSideboard BoardType = "sideboard"
	BoardMaybeboard BoardType = "maybeboard" // "Considering" cards
)

type FinishType string

const (
	FinishDefault FinishType = "default"
	FinishFoil    FinishType = "foil"
	FinishNonfoil FinishType = "nonfoil"
	FinishEtched  FinishType = "etched"
)
