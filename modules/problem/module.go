package problem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *ProblemController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewProblemService(db)
	m.Controller = NewProblemController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&Problem{},
	}
}
