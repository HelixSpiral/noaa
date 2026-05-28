package ndbc

import (
	"fmt"
	"io"
	"net/http"
)

type CameraService struct {
	stationID int
}

func NewCameraService(station int) *CameraService {
	return &CameraService{
		stationID: station,
	}
}

func (r *CameraService) Latest() ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.ndbc.noaa.gov/buoycam.php?station=%d", r.stationID))
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
