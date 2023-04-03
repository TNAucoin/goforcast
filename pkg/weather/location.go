package weather

import (
	"encoding/json"
	"errors"
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

// GetLatLonFromZip Uses GeoCode to reverse zip into lat lon coords
func (l *Location) GetLatLonFromZip(zip string, apiKey string) (*Location, error) {
	url := fmt.Sprintf("%s%s&limit=%d&appid=%s", GeoLocationUrl, zip, 1, apiKey)
	resp, err := client.Get(url)
	if err != nil {
		return l, errors.New(fmt.Sprintf("failed to make location get request: %s", err))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)

		err = decoder.Decode(&l)
		if err != nil {
			return l, errors.New(fmt.Sprintf("Failed to decode location %s", err))
		}
		return l, nil
	}

	return l, errors.New(fmt.Sprintf("location fetch failed with status: %d", resp.StatusCode))
}
