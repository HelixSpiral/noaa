package spc

type Client struct {
	Reports *Service
}

type Service struct {
	Hail    *HailService
	Tornado *TornadoService
	Wind    *WindService
}

// expectedColumns is the expected number of columns for the hail, torando, and wind CSVs
const expectedColumns = 8

const (
	colTime = iota
	colMetric
	colLocation
	colCounty
	colState
	colLatitude
	colLongitude
	colComments
)

func New() *Client {
	return &Client{
		Reports: &Service{
			Hail:    &HailService{},
			Tornado: &TornadoService{},
			Wind:    &WindService{},
		},
	}
}
