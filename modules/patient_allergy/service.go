package patient_allergy

import (
	"saas-api/core"
	"saas-api/modules/allergy_reaction"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type AllergyService struct {
	*core.BaseService[Allergy]
}

func NewAllergyService(db *gorm.DB) *AllergyService {
	service := core.NewBaseService[Allergy](db)

	// Define searchable fields
	service.SearchableFields = []string{
		"patient_id_no",
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

// Create overrides the base Create to handle ReactionIDs mapping
func (s *AllergyService) Create(allergy *Allergy, reactionIDs []uuid.UUID) error {
	if len(reactionIDs) > 0 {
		var reactions []*allergy_reaction.AllergyReaction
		if err := s.DB.Where("id IN ?", reactionIDs).Find(&reactions).Error; err != nil {
			return err
		}
		allergy.Reactions = reactions
	}
	return s.DB.Create(allergy).Error
}

// GetByPatientID retrieves all allergies for a specific patient
func (s *AllergyService) GetByPatientID(patientID string) ([]*Allergy, error) {
	var allergies []*Allergy
	err := s.DB.Where("patient_id_no = ?", patientID).
		Preload("AllergyName").
		Preload("Reactions").
		Find(&allergies).Error
	if err != nil {
		return nil, err
	}
	return allergies, nil
}
