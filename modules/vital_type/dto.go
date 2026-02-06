package vital_type

// CreateVitalTypeDTO defines fields for creating a vital type
type CreateVitalTypeDTO struct {
	Name      string   `json:"name" binding:"required"`
	Unit      string   `json:"unit" binding:"required"`
	NormalMin *float64 `json:"normal_min"`
	NormalMax *float64 `json:"normal_max"`
}

// UpdateVitalTypeDTO defines fields for updating a vital type
type UpdateVitalTypeDTO struct {
	Name      *string  `json:"name,omitempty"`
	Unit      *string  `json:"unit,omitempty"`
	NormalMin *float64 `json:"normal_min"`
	NormalMax *float64 `json:"normal_max"`
}
