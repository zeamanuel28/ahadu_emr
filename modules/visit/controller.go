package visit

import (
	"saas-api/core"
)

type VisitController struct {
	core.BaseController[Visit, CreateVisitDTO, UpdateVisitDTO]
	Service *VisitService
}

func NewVisitController(service *VisitService) *VisitController {
	return &VisitController{
		BaseController: *core.NewBaseController[Visit, CreateVisitDTO, UpdateVisitDTO](service.BaseService),
		Service:        service,
	}
}
