package branch

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *BranchController) {
	group := parentGroup.Group("/branches")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Branch, CreateBranchDTO, UpdateBranchDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Branches",
		PathPrefix:     "/branches",
		CreateSchema:   "branch.CreateBranchDTO",
		UpdateSchema:   "branch.UpdateBranchDTO",
		ResponseSchema: "branch.Branch",
	})
}
