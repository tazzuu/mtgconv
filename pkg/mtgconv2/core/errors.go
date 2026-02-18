package core

import "fmt"

// custom error types used in the package

type UnrecognizedDomain struct {
	Message string
}

func (e *UnrecognizedDomain) Error() string {
	return fmt.Sprintf("unrecognized domain: %s", e.Message)
}

type UnknownInputSource struct {
	Source InputSource
}
func (e *UnknownInputSource) Error() string {
	return fmt.Sprintf("unknown input format: %s", e.Source)
}

type UnknownOutputFormat struct {
	Format OutputFormat
}

func (e *UnknownOutputFormat) Error() string {
	return fmt.Sprintf("unknown output format: %s", e.Format)
}

type InvalidJSONResponse struct {
	Response []byte
}

func (e *InvalidJSONResponse) Error() string {
	// return "response is not valid JSON"
	return fmt.Sprintf("response is not valid JSON: %s", e.Response)
}

type UnexpectedStatus struct {
	Status string
	StatusCode int
}

func (e *UnexpectedStatus) Error() string {
	return fmt.Sprintf("unexpected status: %s", e.Status)
}

type DeckIDParseError struct {
	URL string
}

func (e *DeckIDParseError) Error() string {
	return fmt.Sprintf("could not get Deck ID from URL %v", e.URL)
}

type JSONParseError struct {
	Body string
	Message string
}

func (e *JSONParseError) Error() string {
	return "error parsing JSON body"
}

type UnknownBoardType struct {
	BoardType string
}

func (e *UnknownBoardType) Error() string {
	return fmt.Sprintf("unknown board type: %s", e.BoardType)
}

type InvalidQuantity struct {
	Quantity int
}

func (e *InvalidQuantity) Error() string {
	return fmt.Sprintf("invalid quantity: %d", e.Quantity)
}

type UnknownSortType struct {
	Type string
}

func (e *UnknownSortType) Error() string {
	return fmt.Sprintf("unknown Sort type: %s", e.Type)
}

type UnknownBracket struct {
	Bracket int
}

func (e *UnknownBracket) Error() string {
	return fmt.Sprintf("unknown Bracket value: %d", e.Bracket)
}

type UnknownSortDirection struct {
	Direction string
}

func (e *UnknownSortDirection) Error() string {
	return fmt.Sprintf("unknown sort direction: %s", e.Direction)
}

type UnknownDeckFormat struct {
	Format string
}

func (e *UnknownDeckFormat) Error() string {
	return fmt.Sprintf("unknown deck format: %s", e.Format)
}

type UnknownFinishType struct {
	Finish string
}

func (e *UnknownFinishType) Error() string {
	return fmt.Sprintf("unknown Finish type: %s", e.Finish)
}

type TemplateInitializationError struct {
	Message error
}

func (e *TemplateInitializationError) Error() string {
	return fmt.Sprintf("error initializing template: %v", e.Message)
}

type TemplateExecutionError struct {
	Message error
}

func (e *TemplateExecutionError) Error() string {
	return fmt.Sprintf("error executing template: %v", e.Message)
}

type LineParseError struct {
	Message string
}
func (e *LineParseError) Error() string {
	return fmt.Sprintf("error parsing line: %s", e.Message)
}

type QuantityParseError struct {
	Quantity string
}
func (e *QuantityParseError) Error() string {
	return fmt.Sprintf("error parsing quantity: %s", e.Quantity)
}


type InvalidToken struct {}
func (e *InvalidToken) Error() string {
	return "Invalid API token supplied for user agent"
}