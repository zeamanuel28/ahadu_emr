package observation

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	ObsController       *ObservationController
	ComplaintController *ChiefComplaintController
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Init(db *gorm.DB) {
	service := NewObservationService(db)
	m.ObsController = NewObservationController(service)
	m.ComplaintController = NewChiefComplaintController(service)
}

func (m *Module) RegisterRoutes(parentGroup *gin.RouterGroup) {
	RegisterRoutes(parentGroup, m.ObsController, m.ComplaintController)
}

func (m *Module) GetModels() []interface{} {
	return []interface{}{
		&ChiefComplaint{},
		&Observation{},
	}
}
