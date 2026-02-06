package vital_record

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *VitalRecordController) {
	group := parentGroup.Group("/vital-records")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[VitalRecord, CreateVitalRecordDTO, UpdateVitalRecordDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Vital Records",
		PathPrefix:     "/vital-records",
		CreateSchema:   "vital_record.CreateVitalRecordDTO",
		UpdateSchema:   "vital_record.UpdateVitalRecordDTO",
		ResponseSchema: "vital_record.VitalRecord",
	})
}
