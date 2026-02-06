package visit

import (
	"net/http"
	"saas-api/core"
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
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

// AddObservation adds an observation to a visit
func (ctrl *VisitController) AddObservation(c *gin.Context) {
	var dto CreateObservationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	obs := Observation{
		VisitID:          dto.VisitID,
		ChiefComplaintID: dto.ChiefComplaintID,
		Note:             dto.Note,
	}

	if err := ctrl.Service.CreateObservation(&obs); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create observation", err.Error())
		return
	}

	utils.CreatedResponse(c, obs)
}

// AddVitalRecord adds a vital record to a visit
func (ctrl *VisitController) AddVitalRecord(c *gin.Context) {
	var dto CreateVitalRecordDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	vital := VitalRecord{
		VisitID: dto.VisitID,
		Values:  dto.Values,
	}

	if err := ctrl.Service.CreateVitalRecord(&vital); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create vital record", err.Error())
		return
	}

	utils.CreatedResponse(c, vital)
}
