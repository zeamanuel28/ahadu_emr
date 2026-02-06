package problem

import (
	"errors"
	"saas-api/core"

	"gorm.io/gorm"
)

type ProblemService struct {
	*core.BaseService[Problem]
}

func NewProblemService(db *gorm.DB) *ProblemService {
	service := core.NewBaseService[Problem](db)
	service.Preloads = []string{"DiagnosisCode"}
	return &ProblemService{
		BaseService: service,
	}
}

// Create validates that the patient and diagnosis code exist
func (s *ProblemService) Create(problem *Problem) error {
	// Validate patient existence
	var patientCount int64
	if err := s.DB.Table("patients").Where("id = ? AND is_deleted = ?", problem.PatientID, false).Count(&patientCount).Error; err != nil {
		return err
	}
	if patientCount == 0 {
		return errors.New("patient does not exist")
	}

	// Validate diagnosis code existence
	var codeCount int64
	if err := s.DB.Table("diagnosis_codes").Where("id = ? AND is_deleted = ?", problem.DiagnosisCodeID, false).Count(&codeCount).Error; err != nil {
		return err
	}
	if codeCount == 0 {
		return errors.New("diagnosis code does not exist")
	}

	return s.BaseService.Create(problem)
}
