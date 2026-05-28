package ndbc

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

type Bouy struct {
	ID      int    `json:"id"`
	Picture []byte `json:"picture"`

	MeteorologicalData MeteorologicalData `json:"meteorological_data"`
}

type MeteorologicalData struct {
	Timestamp     time.Time `json:"timestamp"`
	WindDirection string    `json:"wind_direction"`
	WindSpeed     float64   `json:"wind_speed"`
	GustSpeed     float64   `json:"gust_speed"`
}

func (n *Client) GetPictureFromBuoy(id int) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.ndbc.noaa.gov/buoycam.php?station=%d", id))
	if err != nil {
		return nil, fmt.Errorf("error in http Get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body %w", err)
	}

	return body, nil
}

func (n *Client) GetLatestDataFromBuoy(id int) (MeteorologicalData, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.ndbc.noaa.gov/data/realtime2/%d.txt", id))
	if err != nil {
		return MeteorologicalData{}, fmt.Errorf("error in http Get: %w", err)
	}
	defer resp.Body.Close()

	csvReader := csv.NewReader(resp.Body)

	records, err := csvReader.ReadAll()
	if err != nil {
		return MeteorologicalData{}, fmt.Errorf("error reading CSV records: %w", err)
	}

	md, err := createStructFromCSV(records[3])
	if err != nil {
		return MeteorologicalData{}, fmt.Errorf("error creating struct: %w", err)
	}

	return md, nil
}
