package noaa

import "time"

type Client struct {
	NDBC *NationalDataBuoyCenter
}

type NationalDataBuoyCenter struct{}

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
