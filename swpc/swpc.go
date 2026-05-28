package swpc

import "github.com/helixspiral/noaa/swpc/solarwind"

type Client struct {
	SolarWind *solarwind.Service
}

func New() *Client {
	return &Client{
		SolarWind: solarwind.New(),
	}
}
