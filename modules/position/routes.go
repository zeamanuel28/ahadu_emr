package position

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *PositionController) {
	group := parentGroup.Group("/positions")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Position, CreatePositionDTO, UpdatePositionDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Positions",
		PathPrefix:     "/positions",
		CreateSchema:   "position.CreatePositionDTO",
		UpdateSchema:   "position.UpdatePositionDTO",
		ResponseSchema: "position.Position",
	})
}
