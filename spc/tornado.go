package spc

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
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
}

func (s *TornadoService) GetTornadoReports(date time.Time) ([]TornadoReport, error) {
	var reports []TornadoReport

	queryUrl := fmt.Sprintf("https://www.spc.noaa.gov/climo/reports/%s_rpts_torn.csv", date.Format("060102"))

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
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println("error processing tornado report:", err)

			continue
		}
		report := TornadoReport{
			Location: row[2],
			County:   row[3],
			State:    row[4],
			Comments: row[7],
		}

		report.Time, err = strconv.Atoi(row[0])
		if err != nil {
			log.Println("error processing tornado report:", err)

			continue
		}
		report.F_Scale, err = parseInt(row[1])
		if err != nil {
			log.Println("error processing tornado report:", err)

			continue
		}
		report.Latitude, err = strconv.ParseFloat(row[5], 64)
		if err != nil {
			log.Println("error processing tornado report:", err)

			continue
		}
		report.Longitude, err = strconv.ParseFloat(row[6], 64)
		if err != nil {
			log.Println("error processing tornado report:", err)

			continue
		}

		reports = append(reports, report)
	}

	return reports, nil
}
