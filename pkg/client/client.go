package client

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{Timeout: 10 * time.Second}
}

// Get returns a http.Response for the given URL
func Get(url string) (*http.Response, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return &http.Response{}, errors.New(fmt.Sprintf("failed to get: %s", err))
	}
	return resp, nil
}
