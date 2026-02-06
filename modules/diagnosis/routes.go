package diagnosis

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *DiagnosisController) {
	group := parentGroup.Group("/diagnoses")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[Diagnosis, CreateDiagnosisDTO, UpdateDiagnosisDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Diagnoses",
		PathPrefix:     "/diagnoses",
		CreateSchema:   "diagnosis.CreateDiagnosisDTO",
		UpdateSchema:   "diagnosis.UpdateDiagnosisDTO",
		ResponseSchema: "diagnosis.Diagnosis",
	})
}
