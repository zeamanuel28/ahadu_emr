package observation

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, obsCtrl *ObservationController, complaintCtrl *ChiefComplaintController) {
	// Observation Routes
	obsGroup := parentGroup.Group("/observations")
	obsGroup.Use(middleware.AuthMiddleware())

	baseRouteObs := core.NewBaseRoute[Observation, CreateObservationDTO, UpdateObservationDTO](obsCtrl)
	baseRouteObs.Register(obsGroup, core.RouteOptions{
		Tag:            "Observations",
		PathPrefix:     "/observations",
		CreateSchema:   "observation.CreateObservationDTO",
		UpdateSchema:   "observation.UpdateObservationDTO",
		ResponseSchema: "observation.Observation",
	})

	// Chief Complaint Routes
	complaintGroup := parentGroup.Group("/chief-complaints")
	complaintGroup.Use(middleware.AuthMiddleware())

	baseRouteComplaint := core.NewBaseRoute[ChiefComplaint, CreateChiefComplaintDTO, UpdateChiefComplaintDTO](complaintCtrl)
	baseRouteComplaint.Register(complaintGroup, core.RouteOptions{
		Tag:            "Chief Complaints",
		PathPrefix:     "/chief-complaints",
		CreateSchema:   "observation.CreateChiefComplaintDTO",
		UpdateSchema:   "observation.UpdateChiefComplaintDTO",
		ResponseSchema: "observation.ChiefComplaint",
	})
}
