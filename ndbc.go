package noaa

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

func (n *NationalDataBuoyCenter) GetPictureFromBuoy(id int) ([]byte, error) {
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

func (n *NationalDataBuoyCenter) GetLatestDataFromBuoy(id int) (MeteorologicalData, error) {
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
