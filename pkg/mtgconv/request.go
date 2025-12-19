package mtgconv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
	"context"

	"golang.org/x/time/rate"
)

// API rate limit 1 query per second
var apiLimiter = rate.NewLimiter(rate.Every(time.Second), 1)


func MakeAPIUrl(deckID string) string {
	u, _ := url.Parse(apiBaseUrl)
	u.Path = path.Join(u.Path, deckID)
	return (u.String())
}

func FetchJSON(ctx context.Context, url string, userAgent string) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if err := apiLimiter.Wait(ctx); err != nil {
		return "", err
	}

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
