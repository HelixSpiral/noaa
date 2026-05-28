package noaa

import (
	"github.com/helixspiral/noaa/ndbc"
	"github.com/helixspiral/noaa/swpc"
)

func New() *Client {
	return &Client{
		NDBC: ndbc.New(),
		SWPC: swpc.New(),
	}
}

type Client struct {
	NDBC *ndbc.Client
	SWPC *swpc.Client
}
