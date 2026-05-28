package solarwind

type PlasmaService struct{}

// PlasmaObservation is a struct that represents the data returned by the NOAA Solar Wind Prediction Center when calling rtsw_wind_1m
type PlasmaObservation struct {
	TimeTag            string  `json:"time_tag"`
	Active             bool    `json:"active"`
	Source             string  `json:"source"`
	ProtonSpeed        float64 `json:"proton_speed"`
	ProtonTemperature  int     `json:"proton_temperature"`
	ProtonDensity      float64 `json:"proton_density"`
	ProtonVxGse        any     `json:"proton_vx_gse"`
	ProtonVyGse        any     `json:"proton_vy_gse"`
	ProtonVzGse        any     `json:"proton_vz_gse"`
	ProtonVxGsm        any     `json:"proton_vx_gsm"`
	ProtonVyGsm        any     `json:"proton_vy_gsm"`
	ProtonVzGsm        any     `json:"proton_vz_gsm"`
	ProtonSampleSize   int     `json:"proton_sample_size"`
	AlphaSpeed         any     `json:"alpha_speed"`
	AlphaTemperature   any     `json:"alpha_temperature"`
	AlphaDensity       any     `json:"alpha_density"`
	AlphaVxGse         any     `json:"alpha_vx_gse"`
	AlphaVyGse         any     `json:"alpha_vy_gse"`
	AlphaVzGse         any     `json:"alpha_vz_gse"`
	AlphaVxGsm         any     `json:"alpha_vx_gsm"`
	AlphaVyGsm         any     `json:"alpha_vy_gsm"`
	AlphaVzGsm         any     `json:"alpha_vz_gsm"`
	AlphaSampleSize    any     `json:"alpha_sample_size"`
	MaxConvergenceFlag int     `json:"max_convergence_flag"`
	MaxDataFlag        int     `json:"max_data_flag"`
	MaxErrorCountFlag  int     `json:"max_error_count_flag"`
	MaxProcessingFlag  int     `json:"max_processing_flag"`
	MaxRangeFlag       int     `json:"max_range_flag"`
	MaxSampleCountFlag int     `json:"max_sample_count_flag"`
	MaxTelemetryFlag   int     `json:"max_telemetry_flag"`
	OverallQuality     int     `json:"overall_quality"`
}

func (s *PlasmaService) OneMinute() ([]PlasmaObservation, error) {
	rtswURL := "https://services.swpc.noaa.gov/json/rtsw/rtsw_wind_1m.json"

	var po []PlasmaObservation
	err := rawRequest(rtswURL, &po)
	if err != nil {
		return nil, err
	}

	return po, nil
}
