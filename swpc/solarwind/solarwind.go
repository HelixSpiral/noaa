package solarwind

type Service struct {
	Plasma        *PlasmaService
	MagneticField *MagneticFieldService
}

func New() *Service {
	return &Service{
		Plasma:        &PlasmaService{},
		MagneticField: &MagneticFieldService{},
	}
}
