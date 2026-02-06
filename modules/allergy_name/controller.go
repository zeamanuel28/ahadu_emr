package allergy_name

import (
	"saas-api/core"
)

type AllergyNameController struct {
	core.BaseController[AllergyName, CreateAllergyNameDTO, UpdateAllergyNameDTO]
}

func NewAllergyNameController(service *AllergyNameService) *AllergyNameController {
	return &AllergyNameController{
		BaseController: *core.NewBaseController[AllergyName, CreateAllergyNameDTO, UpdateAllergyNameDTO](service.BaseService),
	}
}
