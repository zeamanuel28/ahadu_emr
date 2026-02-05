package allergy

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type AllergyService struct {
	*core.BaseService[Allergy]
}

func NewAllergyService(db *gorm.DB) *AllergyService {
	service := core.NewBaseService[Allergy](db)

	// Define searchable fields for the allergy model
	service.SearchableFields = []string{
		"patient_id",
		"severity",
		"notes",
	}

	// Preload relationships
	service.Preloads = []string{
		"AllergyName",
		"Reactions",
	}

	return &AllergyService{
		BaseService: service,
	}
}

// Custom methods for allergy-specific queries

// GetByPatientID retrieves all allergies for a specific patient
func (s *AllergyService) GetByPatientID(patientID string) ([]*Allergy, error) {
	var allergies []*Allergy
	err := s.DB.Where("patient_id = ?", patientID).
		Preload("AllergyName").
		Preload("Reactions").
		Find(&allergies).Error
	if err != nil {
		return nil, err
	}
	return allergies, nil
}

// GetBySeverity retrieves allergies filtered by severity level
func (s *AllergyService) GetBySeverity(severity string) ([]*Allergy, error) {
	var allergies []*Allergy
	err := s.DB.Where("severity = ?", severity).
		Preload("AllergyName").
		Preload("Reactions").
		Find(&allergies).Error
	if err != nil {
		return nil, err
	}
	return allergies, nil
}
