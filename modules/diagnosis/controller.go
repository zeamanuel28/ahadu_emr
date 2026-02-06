package diagnosis

import (
	"saas-api/core"
)

type DiagnosisController struct {
	core.BaseController[Diagnosis, CreateDiagnosisDTO, UpdateDiagnosisDTO]
}

func NewDiagnosisController(service *DiagnosisService) *DiagnosisController {
	return &DiagnosisController{
		BaseController: *core.NewBaseController[Diagnosis, CreateDiagnosisDTO, UpdateDiagnosisDTO](service.BaseService),
	}
}
