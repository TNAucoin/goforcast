package main

import (
	"github.com/tnaucoin/goforcast/cmd"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {

	cmd.Execute()
}
