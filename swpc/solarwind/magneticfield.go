package solarwind

type MagneticFieldService struct{}

// MagneticObservation is a struct that represents the data returned by the NOAA Solar Wind Prediction Center when calling rtsw_mag_1m
type MagneticObservation struct {
	TimeTag          string  `json:"time_tag"`
	Active           bool    `json:"active"`
	Source           string  `json:"source"`
	Range            any     `json:"range"`
	Scale            any     `json:"scale"`
	Sensitivity      any     `json:"sensitivity"`
	ManualMode       bool    `json:"manual_mode"`
	SampleSize       int     `json:"sample_size"`
	Bt               float64 `json:"bt"`
	BxGse            float64 `json:"bx_gse"`
	ByGse            float64 `json:"by_gse"`
	BzGse            float64 `json:"bz_gse"`
	ThetaGse         float64 `json:"theta_gse"`
	PhiGse           float64 `json:"phi_gse"`
	BxGsm            float64 `json:"bx_gsm"`
	ByGsm            float64 `json:"by_gsm"`
	BzGsm            float64 `json:"bz_gsm"`
	ThetaGsm         float64 `json:"theta_gsm"`
	PhiGsm           float64 `json:"phi_gsm"`
	MaxTelemetryFlag int     `json:"max_telemetry_flag"`
	MaxDataFlag      int     `json:"max_data_flag"`
	OverallQuality   int     `json:"overall_quality"`
}

func (s *MagneticFieldService) OneMinute() ([]MagneticObservation, error) {
	rtswURL := "https://services.swpc.noaa.gov/json/rtsw/rtsw_mag_1m.json"

	var mo []MagneticObservation
	err := rawRequest(rtswURL, &mo)
	if err != nil {
		return nil, err
	}

	return mo, nil
}
