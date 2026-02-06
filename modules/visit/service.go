package visit

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type VisitService struct {
	*core.BaseService[Visit]
}

func NewVisitService(db *gorm.DB) *VisitService {
	service := core.NewBaseService[Visit](db)

	// Define searchable fields
	service.SearchableFields = []string{
		"patient_id_no",
	}

	// Default preloads
	service.Preloads = []string{"Observations", "Vitals"}

	return &VisitService{
		BaseService: service,
	}
}

// Custom methods for Observation and VitalRecord if needed
func (s *VisitService) CreateObservation(obs *Observation) error {
	return s.DB.Create(obs).Error
}

func (s *VisitService) CreateVitalRecord(vital *VitalRecord) error {
	return s.DB.Create(vital).Error
}
