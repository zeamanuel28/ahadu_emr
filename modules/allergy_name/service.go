package allergy_name

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type AllergyNameService struct {
	*core.BaseService[AllergyName]
}

func NewAllergyNameService(db *gorm.DB) *AllergyNameService {
	service := core.NewBaseService[AllergyName](db)

	service.SearchableFields = []string{
		"name",
		"description",
	}

	return &AllergyNameService{
		BaseService: service,
	}
}
