package disposition

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *DispositionController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewDispositionService(db)
	m.Controller = NewDispositionController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&Disposition{},
	}
}
