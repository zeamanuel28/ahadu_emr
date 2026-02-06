package diagnosis_code

import (
	"saas-api/core"

	"gorm.io/gorm"
)

type DiagnosisCodeService struct {
	*core.BaseService[DiagnosisCode]
}

func NewDiagnosisCodeService(db *gorm.DB) *DiagnosisCodeService {
	service := core.NewBaseService[DiagnosisCode](db)
	service.SearchableFields = []string{"code", "name", "description"}
	return &DiagnosisCodeService{
		BaseService: service,
	}
}
