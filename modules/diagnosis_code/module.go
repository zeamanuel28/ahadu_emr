package diagnosis_code

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Controller *DiagnosisCodeController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewDiagnosisCodeService(db)
	m.Controller = NewDiagnosisCodeController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.Controller)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&DiagnosisCode{},
	}
}
