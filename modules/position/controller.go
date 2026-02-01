package position

import (
	"saas-api/core"
)

type PositionController struct {
	core.BaseController[Position, CreatePositionDTO, UpdatePositionDTO]
}

func NewPositionController(service *PositionService) *PositionController {
	return &PositionController{
		BaseController: *core.NewBaseController[Position, CreatePositionDTO, UpdatePositionDTO](&service.BaseService),
	}
}
