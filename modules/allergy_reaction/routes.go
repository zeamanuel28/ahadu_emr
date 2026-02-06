package allergy_reaction

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *AllergyReactionController) {
	group := parentGroup.Group("/allergy-reactions")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[AllergyReaction, CreateAllergyReactionDTO, UpdateAllergyReactionDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Allergy Reactions",
		PathPrefix:     "/allergy-reactions",
		CreateSchema:   "allergy_reaction.CreateAllergyReactionDTO",
		UpdateSchema:   "allergy_reaction.UpdateAllergyReactionDTO",
		ResponseSchema: "allergy_reaction.AllergyReaction",
	})
}
