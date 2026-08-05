package spc

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type HailService struct{}

type HailReport struct {
	Time      int     `json:"time"`
	Size      *int    `json:"size"`
	Location  string  `json:"location"`
	County    string  `json:"county"`
	State     string  `json:"state"`
	Comments  string  `json:"comments"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Errors []error `json:"errors"`
}

func (s *HailService) ByDate(date time.Time) ([]HailReport, error) {
	var reports []HailReport

	queryUrl := fmt.Sprintf("https://www.spc.noaa.gov/climo/reports/%s_rpts_hail.csv", date.Format("060102"))

	resp, err := http.Get(queryUrl)
	if err != nil {
		return reports, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	_, err = reader.Read() // Get rid of header
	if err != nil {
		return reports, err
	}

	for {
		var report HailReport

		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			report.Errors = append(report.Errors, fmt.Errorf("error parsing row: %w", err))
			reports = append(reports, report)

			continue
		}
		report = HailReport{
			Location: row[2],
			County:   row[3],
			State:    row[4],
			Comments: row[7],
		}

		report.Time, err = strconv.Atoi(row[0])
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report time: %w", err))
		}
		report.Size, err = parseInt(row[1])
		if err != nil {
			report.Errors = append(report.Errors, fmt.Errorf("error processing report size: %w", err))
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
