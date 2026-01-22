package mtgconv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"context"
)



// generic method for making the API request and getting a JSON
func RequestJSON(ctx context.Context, url string, userAgent string) (string, error) {
	// TODO: fill this in with the appropriate context object type
	if ctx == nil {
		ctx = context.Background()
	}

	// first check which API domain we are requesting from
	// and update the request object appropriately
	if err := MoxfieldAPIRateLimiter.Wait(ctx); err != nil {
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
