package problem

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *ProblemController) {
	group := parentGroup.Group("/problems")
	group.Use(middleware.AuthMiddleware())

	baseRoute := core.NewBaseRoute[Problem, CreateProblemDTO, UpdateProblemDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Problems",
		PathPrefix:     "/problems",
		CreateSchema:   "problem.CreateProblemDTO",
		UpdateSchema:   "problem.UpdateProblemDTO",
		ResponseSchema: "problem.Problem",
	})
}
