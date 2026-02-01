package employee

import (
	"saas-api/core"
)

type EmployeeController struct {
	core.BaseController[Employee, CreateEmployeeDTO, UpdateEmployeeDTO]
}

func NewEmployeeController(service *EmployeeService) *EmployeeController {
	return &EmployeeController{
		BaseController: *core.NewBaseController[Employee, CreateEmployeeDTO, UpdateEmployeeDTO](&service.BaseService),
	}
}
