package config

import (
	"net/http"
	"time"
)

type ClientConfig struct {
	Timeout    time.Duration
	HttpClient *http.Client
}
