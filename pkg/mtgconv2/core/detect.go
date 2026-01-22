package core

import (
	"log/slog"
	"net/url"
)

func DetectURLSource(urlStr string) (APISource, error) {
	slog.Debug("parsing domain for URL", "urlStr", urlStr)

	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// TODO: add hostname normalization for edge-cases from the parsing of the URL
	hostname := u.Hostname()
	slog.Debug("got domain", "hostname", hostname)

	switch APISource(hostname) {
	case SourceMoxfield:
		return SourceMoxfield, nil
	// add more cases here
	default:
		return "", &UnrecognizedDomain{Message: hostname}
	}
}
