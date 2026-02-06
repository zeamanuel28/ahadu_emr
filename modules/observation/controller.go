package observation

import (
	"saas-api/core"
)

type ObservationController struct {
	core.BaseController[Observation, CreateObservationDTO, UpdateObservationDTO]
}

func NewObservationController(service *ObservationService) *ObservationController {
	return &ObservationController{
		BaseController: *core.NewBaseController[Observation, CreateObservationDTO, UpdateObservationDTO](service.BaseService),
	}
}

type ChiefComplaintController struct {
	core.BaseController[ChiefComplaint, CreateChiefComplaintDTO, UpdateChiefComplaintDTO]
}

func NewChiefComplaintController(service *ObservationService) *ChiefComplaintController {
	return &ChiefComplaintController{
		BaseController: *core.NewBaseController[ChiefComplaint, CreateChiefComplaintDTO, UpdateChiefComplaintDTO](service.ComplaintService),
	}
}
