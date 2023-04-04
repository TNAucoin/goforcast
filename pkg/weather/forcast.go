package weather

import (
	"fmt"
	"github.com/tnaucoin/goforcast/pkg/client"
	"io"
	"log"
)

const currentWeatherURL = "https://api.openweathermap.org/data/2.5/weather"

type Forecast struct {
	Weather Weather     `json:"weather"`
	Temp    Temperature `json:"main"`
}

type Weather struct {
	WeatherType string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Temperature struct {
	Current   float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	Min       float32 `json:"temp_min"`
	Max       float32 `json:"temp_max"`
	Humidity  float32 `json:"humidity"`
}

func GetCurrentWeather(lat, lon float32, apiKey string) (*Forecast, error) {
	url := buildURLRequest(lat, lon, apiKey)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read body res")
	}
	fmt.Printf("%s", body)
	return nil, fmt.Errorf("not implemented")
}

func buildURLRequest(lat, lon float32, apiKey string) string {
	return fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", currentWeatherURL, lat, lon, apiKey)
}
