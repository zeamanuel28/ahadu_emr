package allergy

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *AllergyController) {
	group := parentGroup.Group("/allergies")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Allergy, CreateAllergyDTO, UpdateAllergyDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Allergies",
		PathPrefix:     "/allergies",
		CreateSchema:   "allergy.CreateAllergyDTO",
		UpdateSchema:   "allergy.UpdateAllergyDTO",
		ResponseSchema: "allergy.Allergy",
	})
}
