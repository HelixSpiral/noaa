package noaa

func New() *Client {
	return &Client{
		NDBC: &NationalDataBuoyCenter{},
	}
}
