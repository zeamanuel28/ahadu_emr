package employee

import (
	"saas-api/core"
	"saas-api/shared/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(parentGroup *gin.RouterGroup, controller *EmployeeController) {
	group := parentGroup.Group("/employees")
	group.Use(middleware.AuthMiddleware())
	baseRoute := core.NewBaseRoute[Employee, CreateEmployeeDTO, UpdateEmployeeDTO](controller)
	baseRoute.Register(group, core.RouteOptions{
		Tag:            "Employees",
		PathPrefix:     "/employees",
		CreateSchema:   "employee.CreateEmployeeDTO",
		UpdateSchema:   "employee.UpdateEmployeeDTO",
		ResponseSchema: "employee.Employee",
	})
}
