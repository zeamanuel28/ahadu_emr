package diagnosis_code

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *DiagnosisCodeController) {
	group := parentGroup.Group("/diagnosis-codes")
	group.Use(middleware.AuthMiddleware())

	// Register base CRUD routes
	baseRoute := core.NewBaseRoute[DiagnosisCode, CreateDiagnosisCodeDTO, UpdateDiagnosisCodeDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Diagnosis Codes",
		PathPrefix:     "/diagnosis-codes",
		CreateSchema:   "diagnosis_code.CreateDiagnosisCodeDTO",
		UpdateSchema:   "diagnosis_code.UpdateDiagnosisCodeDTO",
		ResponseSchema: "diagnosis_code.DiagnosisCode",
	})
}
