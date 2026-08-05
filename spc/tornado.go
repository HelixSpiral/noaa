package spc

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

type TornadoService struct{}

type TornadoReport struct {
	Time      int     `json:"time"`
	F_Scale   *int    `json:"f_scale"`
	Location  string  `json:"location"`
	County    string  `json:"county"`
	State     string  `json:"state"`
	Comments  string  `json:"comments"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Errors []error `json:"errors"`
}

func (s *TornadoService) GetTornadoReports(date time.Time) ([]TornadoReport, error) {
	var reports []TornadoReport

	queryURL := fmt.Sprintf("https://www.spc.noaa.gov/climo/reports/%s_rpts_torn.csv", date.Format("060102"))

	body, err := fetch(queryURL)
	if err != nil {
		return reports, err
	}
	defer body.Close()

	reader := csv.NewReader(body)

	_, err = reader.Read() // Get rid of header
	if err != nil {
		return reports, err
	}

	for {
		var report TornadoReport

		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			report.Errors = append(report.Errors, fmt.Errorf("error parsing row: %w", err))
			reports = append(reports, report)

			continue
		}
		report = TornadoReport{
			Location: row[2],
			County:   row[3],
			State:    row[4],
			Comments: row[7],
		}

		report.Time, err = strconv.Atoi(row[0])
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report time: %w", err))
		}
		report.F_Scale, err = parseInt(row[1])
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report f scale: %w", err))
		}
		report.Latitude, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report latitude: %w", err))
		}
		report.Longitude, err = strconv.ParseFloat(row[6], 64)
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report longitude: %w", err))
		}

		reports = append(reports, report)
	}

	return reports, nil
}
