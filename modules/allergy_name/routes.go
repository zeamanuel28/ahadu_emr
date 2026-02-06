package allergy_name

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *AllergyNameController) {
	group := parentGroup.Group("/allergy-names")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[AllergyName, CreateAllergyNameDTO, UpdateAllergyNameDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Allergy Names",
		PathPrefix:     "/allergy-names",
		CreateSchema:   "allergy_name.CreateAllergyNameDTO",
		UpdateSchema:   "allergy_name.UpdateAllergyNameDTO",
		ResponseSchema: "allergy_name.AllergyName",
	})
}
