package allergy_reaction

import (
	"saas-api/core"
)

type AllergyReactionController struct {
	core.BaseController[AllergyReaction, CreateAllergyReactionDTO, UpdateAllergyReactionDTO]
}

func NewAllergyReactionController(service *AllergyReactionService) *AllergyReactionController {
	return &AllergyReactionController{
		BaseController: *core.NewBaseController[AllergyReaction, CreateAllergyReactionDTO, UpdateAllergyReactionDTO](service.BaseService),
	}
}
