package vital_record

import (
	"encoding/json"
	"errors"
	"fmt"
	"saas-api/core"
	"saas-api/modules/vital_type"

	"gorm.io/gorm"
)

type VitalRecordService struct {
	*core.BaseService[VitalRecord]
}

func NewVitalRecordService(db *gorm.DB) *VitalRecordService {
	service := core.NewBaseService[VitalRecord](db)
	return &VitalRecordService{
		BaseService: service,
	}
}

// Create validates that all vital types in the JSON values exist and that the visit exists
func (s *VitalRecordService) Create(record *VitalRecord) error {
	// Validate that the visit exists before creating the vital record
	var count int64
	err := s.DB.Table("visits").Where("id = ? AND is_deleted = ?", record.VisitID, false).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("visit with ID " + record.VisitID.String() + " does not exist")
	}

	// Parse the JSON values to get keys
	var values map[string]interface{}
	data, _ := record.Values.MarshalJSON()
	if err := json.Unmarshal(data, &values); err != nil {
		return errors.New("invalid JSON format for vital values")
	}

	// For each key in the JSON, verify it exists as a VitalType Name and validate values against min/max
	for vitalName, value := range values {
		// Check if the vital type exists
		var vitalType vital_type.VitalType
		err := s.DB.Where("name = ? AND is_deleted = ?", vitalName, false).First(&vitalType).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("vital type '%s' is not defined in the system", vitalName)
			}
			return err
		}

		// Validate the value against min/max if they are defined
		if vitalType.NormalMin != nil || vitalType.NormalMax != nil {
			// Convert the value to float64 for comparison
			floatValue, ok := convertToFloat64(value)
			if !ok {
				return fmt.Errorf("vital value for '%s' must be a number to compare against min/max ranges", vitalName)
			}

			// Check if value is below minimum
			if vitalType.NormalMin != nil && floatValue < *vitalType.NormalMin {
				return fmt.Errorf("vital value for '%s' (%f) is below normal minimum (%f)", vitalName, floatValue, *vitalType.NormalMin)
			}

			// Check if value is above maximum
			if vitalType.NormalMax != nil && floatValue > *vitalType.NormalMax {
				return fmt.Errorf("vital value for '%s' (%f) is above normal maximum (%f)", vitalName, floatValue, *vitalType.NormalMax)
			}
		}
	}

	return s.BaseService.Create(record)
}

// Helper function to convert interface{} to float64
func convertToFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int32:
		return float64(v), true
	case int64:
		return float64(v), true
	case string:
		// Attempt to parse string as float
		var result float64
		_, err := fmt.Sscanf(v, "%f", &result)
		if err == nil {
			return result, true
		}
		return 0, false
	default:
		return 0, false
	}
}
