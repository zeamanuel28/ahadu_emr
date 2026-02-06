package allergy_reaction

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type AllergyReactionService struct {
	*core.BaseService[AllergyReaction]
}

func NewAllergyReactionService(db *gorm.DB) *AllergyReactionService {
	service := core.NewBaseService[AllergyReaction](db)

	service.SearchableFields = []string{
		"name",
		"description",
	}

	return &AllergyReactionService{
		BaseService: service,
	}
}
