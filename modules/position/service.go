package position

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type PositionService struct {
	core.BaseService[Position]
}

func NewPositionService(db *gorm.DB) *PositionService {
	return &PositionService{
		BaseService: *core.NewBaseService[Position](db),
	}
}
