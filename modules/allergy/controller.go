package allergy

import (
	"saas-api/core"
)

type AllergyController struct {
	core.BaseController[Allergy, CreateAllergyDTO, UpdateAllergyDTO]
}

func NewAllergyController(service *AllergyService) *AllergyController {
	return &AllergyController{
		BaseController: *core.NewBaseController[Allergy, CreateAllergyDTO, UpdateAllergyDTO](service.BaseService),
	}
}
