package vital_type

import (
	"saas-api/core"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, ctrl *VitalTypeController) {
	group := parentGroup.Group("/vital-types")

	baseRoute := core.NewBaseRoute[VitalType, CreateVitalTypeDTO, UpdateVitalTypeDTO](ctrl)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "VitalTypes",
		PathPrefix:     "/vital-types",
		CreateSchema:   "vital_type.CreateVitalTypeDTO",
		UpdateSchema:   "vital_type.UpdateVitalTypeDTO",
		ResponseSchema: "vital_type.VitalType",
	})
}
