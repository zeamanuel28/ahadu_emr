package vital_type

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type VitalTypeService struct {
	*core.BaseService[VitalType]
}

func NewVitalTypeService(db *gorm.DB) *VitalTypeService {
	service := core.NewBaseService[VitalType](db)
	return &VitalTypeService{
		BaseService: service,
	}
}
