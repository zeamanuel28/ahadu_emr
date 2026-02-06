package observation

import (
	"errors"
	"saas-api/core"

	"gorm.io/gorm"
)

type ObservationService struct {
	*core.BaseService[Observation]
	ComplaintService *core.BaseService[ChiefComplaint]
}

func NewObservationService(db *gorm.DB) *ObservationService {
	obsService := core.NewBaseService[Observation](db)
	obsService.Preloads = []string{"ChiefComplaint"}
	obsService.SearchableFields = []string{"note"}

	complaintService := core.NewBaseService[ChiefComplaint](db)
	complaintService.SearchableFields = []string{"name", "description"}

	return &ObservationService{
		BaseService:      obsService,
		ComplaintService: complaintService,
	}
}

// Create validates that the visit exists before creating the observation
func (s *ObservationService) Create(observation *Observation) error {
	// Validate that the visit exists before creating the observation
	var visitCount int64
	err := s.DB.Table("visits").Where("id = ? AND is_deleted = ?", observation.VisitID, false).Count(&visitCount).Error
	if err != nil {
		return err
	}
	if visitCount == 0 {
		return errors.New("visit with ID " + observation.VisitID.String() + " does not exist")
	}

	// Proceed with creating the observation
	return s.BaseService.Create(observation)
}

// Custom methods for ObservationService can be added here
