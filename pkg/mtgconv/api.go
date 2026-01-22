package mtgconv
import (
	"log/slog"
	"net/url"
	"fmt"
)

func CheckConnectivity() error {
	return nil
}

// enum for the various API's we will recognize for loading data from
type APISource string
const (
	MoxfieldAPIType APISource = "moxfield.com"
	// TODO: add more API Sources here ...
	// NOTE: also add any new API Sources to the DetectURLSource
)

func (a APISource) String() string {
	return string(a)
}



// custom error for a domain we dont recognize as per above
type UnrecognizedDomain struct {
	Message string
}
func (e *UnrecognizedDomain) Error() string {
    return fmt.Sprintf("UnrecognizedDomain error: %s", e.Message)
}

// check which website a URL is point to
func DetectURLSource(urlStr string) (APISource, error) {
	slog.Debug("parsing domain for URL", "urlStr", urlStr)
    urlObj, err := url.Parse(urlStr)
    if err != nil {
        slog.Error("error parsing URL", "err", err)
		return "", err
    }

    hostname := urlObj.Hostname()
	slog.Debug("got domain", "hostname", hostname)

	switch APISource(hostname) {
	case MoxfieldAPIType:
		return MoxfieldAPIType, nil
	// add more cases here
	default:
		return "", &UnrecognizedDomain{Message: hostname}
	}
}