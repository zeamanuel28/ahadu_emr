package disposition

import (
	"saas-api/core"
)

type DispositionController struct {
	*core.BaseController[Disposition, CreateDispositionDTO, UpdateDispositionDTO]
}

func NewDispositionController(service *DispositionService) *DispositionController {
	return &DispositionController{
		BaseController: core.NewBaseController[Disposition, CreateDispositionDTO, UpdateDispositionDTO](service.BaseService),
	}
}
