package weather

import (
	"encoding/json"
	"fmt"
	"github.com/tnaucoin/goforcast/pkg/client"
	"net/http"
)

// Base url for making a weather API request
const currentWeatherURL = "https://api.openweathermap.org/data/2.5/weather"

// Forecast struct for the weather API data
type Forecast struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// GetCurrentWeather retrieves the current weather forecast from a given lat lon location
func GetCurrentWeather(lat, lon float32, apiKey string) (*Forecast, error) {
	url := buildURLRequest(lat, lon, apiKey)
	fmt.Println(url)
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather request returned: %d", res.StatusCode)
	}
	decode := json.NewDecoder(res.Body)
	var forecast Forecast
	err = decode.Decode(&forecast)
	if err != nil {
		return nil, err
	}

	return &forecast, nil
}

// buildURLRequest builds the url string with the correct params
func buildURLRequest(lat, lon float32, apiKey string) string {
	return fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", currentWeatherURL, lat, lon, apiKey)
}
