package ndbc

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RealTimeService struct {
	stationID int
}

func NewRealtimeService(station int) *RealTimeService {
	return &RealTimeService{
		stationID: station,
	}
}

type StandardObservation struct {
	Timestamp          time.Time `json:"timestamp"`
	WindDirection      int       `json:"wind_direction"`
	WindSpeed          float64   `json:"wind_speed"`
	GustSpeed          float64   `json:"gust_speed"`
	WaveHeight         float64   `json:"wave_height"`
	DominantWavePeriod float64   `json:"dominant_wave_period"`
	AverageWaevPeriod  float64   `json:"average_wave_period"`
	WaveDirection      int       `json:"wave_direction"`
	Pressure           float64   `json:"pressure"`
	AirTemp            float64   `json:"air_temp"`
	WaterTemp          float64   `json:"water_temp"`
	DewPoint           float64   `json:"dew_point"`
	PressureTendency   float64   `json:"pressure_tendency"`
}

func (r *RealTimeService) Standard() ([]StandardObservation, error) {
	var records []StandardObservation
	resp, err := http.Get(fmt.Sprintf("https://www.ndbc.noaa.gov/data/realtime2/%d.txt", r.stationID))
	if err != nil {
		return records, fmt.Errorf("error in http get: %w", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// Skip headers
		if strings.HasPrefix(line, "#") {
			continue
		}

		row := strings.Fields(line)

		var record StandardObservation

		timestamp, err := time.ParseInLocation("2006/01/02 15:04", fmt.Sprintf("%s/%s/%s %s:%s", row[0], row[1], row[2], row[3], row[4]), time.UTC)
		if err != nil {
			return records, fmt.Errorf("error parsing time: %w", err)
		}

		record.Timestamp = timestamp

		if row[5] != "" && row[5] != "MM" {
			record.WindDirection, err = strconv.Atoi(row[5])
			if err != nil {
				return records, fmt.Errorf("error parsing wind direction: %w", err)
			}
		}

		if row[6] != "" && row[6] != "MM" {
			record.WindSpeed, err = strconv.ParseFloat(row[6], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing wind speed: %w", err)
			}
		}

		if row[7] != "" && row[7] != "MM" {
			record.GustSpeed, err = strconv.ParseFloat(row[7], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing gust speed: %w", err)
			}
		}

		if row[8] != "" && row[8] != "MM" {
			record.WaveHeight, err = strconv.ParseFloat(row[8], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing wave height: %w", err)
			}
		}

		if row[9] != "" && row[9] != "MM" {
			record.DominantWavePeriod, err = strconv.ParseFloat(row[9], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing dominant wave period: %w", err)
			}
		}

		if row[10] != "" && row[10] != "MM" {
			record.AverageWaevPeriod, err = strconv.ParseFloat(row[10], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing average wave period: %w", err)
			}
		}

		if row[11] != "" && row[11] != "MM" {
			record.WaveDirection, err = strconv.Atoi(row[11])
			if err != nil {
				return records, fmt.Errorf("error parsing wave direction: %w", err)
			}
		}

		if row[12] != "" && row[12] != "MM" {
			record.Pressure, err = strconv.ParseFloat(row[12], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing pressure: %w", err)
			}
		}

		if row[13] != "" && row[13] != "MM" {
			record.AirTemp, err = strconv.ParseFloat(row[13], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing air temperature: %w", err)
			}
		}

		if row[14] != "" && row[14] != "MM" {
			record.WaterTemp, err = strconv.ParseFloat(row[14], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing water temperature: %w", err)
			}
		}

		if row[15] != "" && row[15] != "MM" {
			record.DewPoint, err = strconv.ParseFloat(row[15], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing dew point: %w", err)
			}
		}

		if row[17] != "" && row[17] != "MM" {
			record.PressureTendency, err = strconv.ParseFloat(row[17], 64)
			if err != nil {
				return records, fmt.Errorf("error parsing pressure tendency: %w", err)
			}
		}

		records = append(records, record)
	}

	return records, nil
}
