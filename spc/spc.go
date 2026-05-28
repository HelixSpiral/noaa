package spc

type Client struct {
	Reports *Service
}

type Service struct {
	Hail    *HailService
	Tornado *TornadoService
	Wind    *WindService
}

func New() *Client {
	return &Client{
		Reports: &Service{
			Hail:    &HailService{},
			Tornado: &TornadoService{},
			Wind:    &WindService{},
		},
	}
}
