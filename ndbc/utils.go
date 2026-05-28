package ndbc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func createStructFromCSV(data []string) (MeteorologicalData, error) {
	var md MeteorologicalData

	data = strings.Fields(data[0])

	var err error
	md.Timestamp, err = time.ParseInLocation("2006/01/02 15:04", fmt.Sprintf("%s/%s/%s %s:%s", data[0], data[1], data[2], data[3], data[4]), time.UTC)
	if err != nil {
		return MeteorologicalData{}, fmt.Errorf("error parsing time: %w", err)
	}

	for x, y := range data {
		switch x {
		case 5:
			yInt, _ := strconv.Atoi(y)
			md.WindDirection = getWindDirection(yInt)
		case 6:
			yFloat, _ := strconv.ParseFloat(y, 64)
			md.WindSpeed = yFloat
		case 7:
			yFloat, _ := strconv.ParseFloat(y, 64)
			md.GustSpeed = yFloat
		}
	}

	return md, nil
}

func getWindDirection(degrees int) string {
	compas := []string{
		"North", "Northeast",
		"East", "Southeast",
		"South", "Southwest",
		"West", "Northwest",
	}

	return compas[int(math.Round(float64(degrees)/45))%8]
}
