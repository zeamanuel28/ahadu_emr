package department

import (
	"saas-api/core"
)

type DepartmentController struct {
	core.BaseController[Department, CreateDepartmentDTO, UpdateDepartmentDTO]
}

func NewDepartmentController(service *DepartmentService) *DepartmentController {
	return &DepartmentController{
		BaseController: *core.NewBaseController[Department, CreateDepartmentDTO, UpdateDepartmentDTO](&service.BaseService),
	}
}
