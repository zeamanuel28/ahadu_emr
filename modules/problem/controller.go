package problem

import (
	"saas-api/core"
)

type ProblemController struct {
	core.BaseController[Problem, CreateProblemDTO, UpdateProblemDTO]
}

func NewProblemController(service *ProblemService) *ProblemController {
	return &ProblemController{
		BaseController: *core.NewBaseController[Problem, CreateProblemDTO, UpdateProblemDTO](service.BaseService),
	}
}
