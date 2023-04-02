package client

import (
	"fmt"
	"github.com/tnaucoin/goforcast/pkg/config"
	"io"
	"log"
	"net/http"
)

var weatherClient *WeatherClient

type WeatherClient struct {
	Config     *config.ClientConfig
	HttpClient *http.Client
}

func CreateClientWithConfig(clientConfig *config.ClientConfig) *WeatherClient {
	return &WeatherClient{
		Config: clientConfig,
		HttpClient: &http.Client{
			Timeout: clientConfig.Timeout,
		},
	}
}

func NewClient(repoClient *WeatherClient) *WeatherClient {
	weatherClient = repoClient
	return weatherClient
}

func (wc *WeatherClient) Test() {
	resp, err := wc.HttpClient.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(body)
}
