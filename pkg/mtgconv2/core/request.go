package core

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
)

// run a request created by a handler
// and return a JSON response
func DoRequestJSON(req *http.Request)(string, error){
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", &UnexpectedStatus{resp.Status, resp.StatusCode}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if !json.Valid(body) {
		return "", &InvalidJSONResponse{body}
	}

	return string(body), nil
}