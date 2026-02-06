package vital_record

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *VitalRecordController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewVitalRecordService(db)
	m.Controller = NewVitalRecordController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&VitalRecord{},
	}
}
