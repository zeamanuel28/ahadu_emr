package user

import (
	"saas-api/core"
)

type UserController struct {
	core.BaseController[User, CreateUserDTO, UpdateUserDTO]
}

func NewUserController(service *UserService) *UserController {
	return &UserController{
		BaseController: *core.NewBaseController[User, CreateUserDTO, UpdateUserDTO](&service.BaseService),
	}
}
