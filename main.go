package main

import (
	"github.com/tnaucoin/goforcast/cmd"
	"github.com/tnaucoin/goforcast/pkg/client"
	"github.com/tnaucoin/goforcast/pkg/config"
	"time"
)

var httpClientConfig config.ClientConfig

func main() {
	httpClientConfig.Timeout = time.Duration(1) * time.Second
	clientConfig := client.CreateClientWithConfig(&httpClientConfig)
	wc := client.NewClient(clientConfig)
	wc.Test()
	cmd.Execute()
}
