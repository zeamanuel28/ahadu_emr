package user

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *UserController) {
	group := parentGroup.Group("/users")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[User, CreateUserDTO, UpdateUserDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Users",
		PathPrefix:     "/users",
		CreateSchema:   "user.CreateUserDTO",
		UpdateSchema:   "user.UpdateUserDTO",
		ResponseSchema: "user.User",
	})
}
