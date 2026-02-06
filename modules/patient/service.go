package patient

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type PatientService struct {
	*core.BaseService[Patient]
}

func NewPatientService(db *gorm.DB) *PatientService {
	service := core.NewBaseService[Patient](db)

	// Define searchable fields for the patient model
	service.SearchableFields = []string{
		"full_name",
		"id_no",
		"email",
		"phone",
	}

	// Preload allergies and their related data for all patient operations
	service.Preloads = []string{
		"Allergies",
		"Allergies.AllergyName",
		"Allergies.Reactions",
	}

	return &PatientService{
		BaseService: service,
	}
}

// Custom methods can be added here if needed
// For example, if you need specific business logic for patients
func (s *PatientService) GetByPatientID(patientID string) (*Patient, error) {
	var patient Patient
	err := s.DB.Where("id_no = ?", patientID).
		Preload("Allergies").
		Preload("Allergies.AllergyName").
		Preload("Allergies.Reactions").
		First(&patient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

func (s *PatientService) GetByEmail(email string) (*Patient, error) {
	var patient Patient
	err := s.DB.Where("email = ?", email).First(&patient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}

// GetPatientWithAllergies retrieves a patient along with their allergies
func (s *PatientService) GetPatientWithAllergies(patientID string) (*Patient, error) {
	var patient Patient
	err := s.DB.Where("id_no = ?", patientID).
		Preload("Allergies").
		Preload("Allergies.AllergyName").
		Preload("Allergies.Reactions").
		First(&patient).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &patient, nil
}
