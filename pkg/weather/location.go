package weather

import (
	"encoding/json"
	"fmt"
	"github.com/tnaucoin/goforcast/pkg/client"
	"net/http"
)

const geoLocationUrl = "http://api.openweathermap.org/geo/1.0/zip?zip="

// Location stores the lat long and name of the given location
type Location struct {
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
	Name string  `json:"name"`
}

// GetLocationFromZip Uses GeoCode to reverse zip into lat lon coords
func GetLocationFromZip(zip string, apiKey string) (*Location, error) {
	// Construct the URL
	url := fmt.Sprintf("%s%s&limit=%d&appid=%s", geoLocationUrl, zip, 1, apiKey)
	// Send GET req to the API
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make location get request: %s", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("location fetch failed with status: %d", resp.StatusCode)
	}

	// Parse the API response
	decoder := json.NewDecoder(resp.Body)
	var location Location
	err = decoder.Decode(&location)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode location %s", err)
	}
	return &location, nil

}
