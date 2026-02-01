package auth

import (
	"saas-api/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *AuthController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	userService := user.NewUserService(db)
	authService := NewAuthService(userService)
	m.Controller = NewAuthController(authService)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{}
}
