package diagnosis_code

// CreateDiagnosisCodeDTO defines the fields required to create a diagnosis code
type CreateDiagnosisCodeDTO struct {
	Code        string          `json:"code" binding:"required"`
	Name        string          `json:"name" binding:"required"`
	System      DiagnosisSystem `json:"system" binding:"required,oneof=ICD10 SNOMED CUSTOM"`
	Description string          `json:"description,omitempty"`
	Active      *bool           `json:"active,omitempty" binding:"omitempty"`
}

// UpdateDiagnosisCodeDTO defines the fields allowed for updating a diagnosis code
type UpdateDiagnosisCodeDTO struct {
	Code        *string          `json:"code,omitempty"`
	Name        *string          `json:"name,omitempty"`
	System      *DiagnosisSystem `json:"system,omitempty" binding:"omitempty,oneof=ICD10 SNOMED CUSTOM"`
	Description *string          `json:"description,omitempty"`
	Active      *bool            `json:"active,omitempty"`
}
