package patient_allergy

import (
	"net/http"
	"saas-api/core"
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
)

type AllergyController struct {
	core.BaseController[Allergy, CreateAllergyDTO, UpdateAllergyDTO]
	service *AllergyService
}

func NewAllergyController(service *AllergyService) *AllergyController {
	return &AllergyController{
		BaseController: *core.NewBaseController[Allergy, CreateAllergyDTO, UpdateAllergyDTO](service.BaseService),
		service:        service,
	}
}

func (ctrl *AllergyController) Create(c *gin.Context) {
	var dto CreateAllergyDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	allergy := &Allergy{
		PatientID:     dto.PatientID,
		AllergyNameID: dto.AllergyNameID,
		Severity:      dto.Severity,
		Notes:         dto.Notes,
	}

	if err := ctrl.service.Create(allergy, dto.ReactionIDs); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create allergy", err.Error())
		return
	}

	utils.CreatedResponse(c, allergy)
}
