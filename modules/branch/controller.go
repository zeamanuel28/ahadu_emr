package branch

import (
	"saas-api/core"
)

type BranchController struct {
	core.BaseController[Branch, CreateBranchDTO, UpdateBranchDTO]
}

func NewBranchController(service *BranchService) *BranchController {
	return &BranchController{
		BaseController: *core.NewBaseController[Branch, CreateBranchDTO, UpdateBranchDTO](&service.BaseService),
	}
}
