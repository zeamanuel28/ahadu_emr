package patient_allergy

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *AllergyController) {
	group := parentGroup.Group("/patient-allergies")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[Allergy, CreateAllergyDTO, UpdateAllergyDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Patient Allergies",
		PathPrefix:     "/patient-allergies",
		CreateSchema:   "patient_allergy.CreateAllergyDTO",
		UpdateSchema:   "patient_allergy.UpdateAllergyDTO",
		ResponseSchema: "patient_allergy.Allergy",
	})
}
