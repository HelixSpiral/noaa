package ndbc

import (
	"math"
)

func getWindDirection(degrees int) string {
	compas := []string{
		"North", "Northeast",
		"East", "Southeast",
		"South", "Southwest",
		"West", "Northwest",
	}

	return compas[int(math.Round(float64(degrees)/45))%8]
}
