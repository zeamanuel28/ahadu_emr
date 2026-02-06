package visit

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, ctrl *VisitController) {
	group := parentGroup.Group("/visits")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[Visit, CreateVisitDTO, UpdateVisitDTO](ctrl)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Visits",
		PathPrefix:     "/visits",
		CreateSchema:   "visit.CreateVisitDTO",
		UpdateSchema:   "visit.UpdateVisitDTO",
		ResponseSchema: "visit.Visit",
	})

	// Custom endpoints for visit
	group.POST("/observations", ctrl.AddObservation)
	group.POST("/vitals", ctrl.AddVitalRecord)
}
