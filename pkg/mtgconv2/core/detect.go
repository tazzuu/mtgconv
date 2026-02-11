package core

import (
	"log/slog"
	"net/url"
	"strings"
)

// determine which API source was provided based on the input URL
func DetectURLSource(urlStr string) (APISource, InputSource, error) {
	slog.Debug("parsing domain for URL", "urlStr", urlStr)

	if !strings.Contains(urlStr, "://") {
		urlStr = "https://" + urlStr
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return "", "", err
	}

	// TODO: add hostname normalization for edge-cases from the parsing of the URL
	hostname := u.Hostname()
	slog.Debug("got domain", "hostname", hostname)

	api, src, err := ParseAPISource(hostname)
	if err != nil {
		return "", "", err
	}
	return api, src, nil
}
