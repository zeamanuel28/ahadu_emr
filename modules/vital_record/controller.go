package vital_record

import (
	"saas-api/core"
)

type VitalRecordController struct {
	core.BaseController[VitalRecord, CreateVitalRecordDTO, UpdateVitalRecordDTO]
}

func NewVitalRecordController(service *VitalRecordService) *VitalRecordController {
	return &VitalRecordController{
		BaseController: *core.NewBaseController[VitalRecord, CreateVitalRecordDTO, UpdateVitalRecordDTO](service.BaseService),
	}
}
