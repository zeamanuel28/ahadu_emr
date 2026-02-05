package patient

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *PatientController) {
	group := parentGroup.Group("/patients")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Patient, CreatePatientDTO, UpdatePatientDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Patients",
		PathPrefix:     "/patients",
		CreateSchema:   "patient.CreatePatientDTO",
		UpdateSchema:   "patient.UpdatePatientDTO",
		ResponseSchema: "patient.Patient",
	})
}
