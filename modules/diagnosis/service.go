package diagnosis

import (
	"errors"
	"saas-api/core"

	"gorm.io/gorm"
)

type DiagnosisService struct {
	*core.BaseService[Diagnosis]
}

func NewDiagnosisService(db *gorm.DB) *DiagnosisService {
	service := core.NewBaseService[Diagnosis](db)
	service.Preloads = []string{"DiagnosisCode"}
	return &DiagnosisService{
		BaseService: service,
	}
}

// Create validates that the visit and diagnosis code exist
func (s *DiagnosisService) Create(diagnosis *Diagnosis) error {
	// Validate visit existence (using direct DB count to avoid import cycle)
	var visitCount int64
	if err := s.DB.Table("visits").Where("id = ? AND is_deleted = ?", diagnosis.VisitID, false).Count(&visitCount).Error; err != nil {
		return err
	}
	if visitCount == 0 {
		return errors.New("visit does not exist")
	}

	// Validate diagnosis code existence
	var codeCount int64
	if err := s.DB.Table("diagnosis_codes").Where("id = ? AND is_deleted = ?", diagnosis.DiagnosisCodeID, false).Count(&codeCount).Error; err != nil {
		return err
	}
	if codeCount == 0 {
		return errors.New("diagnosis code does not exist")
	}

	return s.BaseService.Create(diagnosis)
}
