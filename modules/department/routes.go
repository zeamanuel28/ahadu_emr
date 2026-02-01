package department

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *DepartmentController) {
	group := parentGroup.Group("/departments")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Department, CreateDepartmentDTO, UpdateDepartmentDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Departments",
		PathPrefix:     "/departments",
		CreateSchema:   "department.CreateDepartmentDTO",
		UpdateSchema:   "department.UpdateDepartmentDTO",
		ResponseSchema: "department.Department",
	})
}
