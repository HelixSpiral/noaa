package noaa

import (
	"github.com/helixspiral/noaa/ndbc"
	"github.com/helixspiral/noaa/spc"
	"github.com/helixspiral/noaa/swpc"
)

type Client struct {
	NDBC *ndbc.Client
	SPC  *spc.Client
	SWPC *swpc.Client
}

func New() *Client {
	return &Client{
		NDBC: ndbc.New(),
		SPC:  spc.New(),
		SWPC: swpc.New(),
	}
}
