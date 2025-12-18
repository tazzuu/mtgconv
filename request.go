package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"path"
)

var apiBaseUrl string = "https://api.moxfield.com/v2/decks/all"

func makeAPIUrl(deckID string) string {
	u, _ := url.Parse(apiBaseUrl)
	u.Path = path.Join(u.Path, deckID)
	return(u.String())
}

func fetchJSON(url string, userAgent string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if !json.Valid(body) {
		return "", errors.New("response is not valid JSON")
	}

	return string(body), nil
}
