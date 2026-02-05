package patient

import (
	"saas-api/core"
)

type PatientController struct {
	core.BaseController[Patient, CreatePatientDTO, UpdatePatientDTO]
}

func NewPatientController(service *PatientService) *PatientController {
	return &PatientController{
		BaseController: *core.NewBaseController[Patient, CreatePatientDTO, UpdatePatientDTO](service.BaseService),
	}
}