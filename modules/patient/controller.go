package patient

import (
	"saas-api/core"
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

type PatientController struct {
	core.BaseController[Patient, CreatePatientDTO, UpdatePatientDTO]
	service *PatientService
}

func NewPatientController(service *PatientService) *PatientController {
	controller := &PatientController{
		BaseController: *core.NewBaseController[Patient, CreatePatientDTO, UpdatePatientDTO](service.BaseService),
		service:        service,
	}
	return controller
}

// GetById overrides the base GetById to search by IDNo (MRN) instead of the default ID
func (ctrl *PatientController) GetById(c *gin.Context) {
	id := c.Param("id")

	// Use the service method that searches by IDNo
	patient, err := ctrl.service.GetByPatientID(id)
	if err != nil {
		utils.ErrorResponse(c, 500, "Failed to fetch patient", err.Error())
		return
	}
	if patient == nil {
		utils.ErrorResponse(c, 404, "Patient not found", "")
		return
	}

	utils.SuccessResponse(c, patient, nil)
}