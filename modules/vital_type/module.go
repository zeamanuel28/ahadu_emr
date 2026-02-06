package vital_type

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *VitalTypeController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewVitalTypeService(db)
	m.Controller = NewVitalTypeController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&VitalType{},
	}
}
