package diagnosis_code

import (
	"saas-api/core"
)

type DiagnosisCodeController struct {
	core.BaseController[DiagnosisCode, CreateDiagnosisCodeDTO, UpdateDiagnosisCodeDTO]
}

func NewDiagnosisCodeController(service *DiagnosisCodeService) *DiagnosisCodeController {
	return &DiagnosisCodeController{
		BaseController: *core.NewBaseController[DiagnosisCode, CreateDiagnosisCodeDTO, UpdateDiagnosisCodeDTO](service.BaseService),
	}
}
