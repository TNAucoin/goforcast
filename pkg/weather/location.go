package weather

import (
	"encoding/json"
	"fmt"
	"github.com/tnaucoin/goforcast/pkg/client"
	"net/http"
)

const GeoLocationUrl = "http://api.openweathermap.org/geo/1.0/zip?zip="

// Location stores the lat long and name of the given location
type Location struct {
	Lat  float32 `json:"lat"`
	Lon  float32 `json:"lon"`
	Name string  `json:"name"`
}

// GetLocationFromZip Uses GeoCode to reverse zip into lat lon coords
func GetLocationFromZip(zip string, apiKey string) (*Location, error) {
	url := fmt.Sprintf("%s%s&limit=%d&appid=%s", GeoLocationUrl, zip, 1, apiKey)
	resp, err := client.Get(url)
	if err != nil {
		return &Location{}, fmt.Errorf("failed to make location get request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		l := Location{}
		err = decoder.Decode(&l)
		if err != nil {
			return &Location{}, fmt.Errorf("Failed to decode location %s", err)
		}
		return &l, nil
	}

	return &Location{}, fmt.Errorf("location fetch failed with status: %d", resp.StatusCode)
}
