package ndbc

type StationService struct {
	RealTime *RealTimeService
	Camera   *CameraService

	id int
}

func (c *Client) Station(id int) *StationService {
	return &StationService{
		id: id,

		Camera:   NewCameraService(id),
		RealTime: NewRealtimeService(id),
	}
}
