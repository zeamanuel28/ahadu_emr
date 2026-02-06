package vital_type

import (
	"saas-api/core"
)

type VitalTypeController struct {
	*core.BaseController[VitalType, CreateVitalTypeDTO, UpdateVitalTypeDTO]
}

func NewVitalTypeController(service *VitalTypeService) *VitalTypeController {
	return &VitalTypeController{
		BaseController: core.NewBaseController[VitalType, CreateVitalTypeDTO, UpdateVitalTypeDTO](service.BaseService),
	}
}
