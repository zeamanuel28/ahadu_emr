package patient

import (
	"saas-api/modules/patient_allergy"
	"time"
)

// CreatePatientDTO defines the fields required to create a patient
type CreatePatientDTO struct {
	PatientImage             string                   `json:"patientImage,omitempty"`
	FullName                 string                   `json:"fullName" binding:"required"`
	Sex                      Sex                      `json:"sex" binding:"required,oneof=MALE FEMALE"`
	DateOfBirth              time.Time                `json:"dateOfBirth" binding:"required"`
	BloodType                BloodType                `json:"bloodType,omitempty" binding:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Phone                    string                   `json:"phone,omitempty" binding:"omitempty,len=10,numeric"`
	Email                    string                   `json:"email,omitempty" binding:"omitempty,email"`
	City                     string                   `json:"city,omitempty"`
	SubCity                  string                   `json:"subcity,omitempty"`
	Wereda                   string                   `json:"wereda,omitempty"`
	EmergencyContactName     string                   `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone    string                   `json:"emergencyContactPhone,omitempty" binding:"omitempty,len=10,numeric"`
	EmergencyContactRelation EmergencyContactRelation `json:"emergencyContactRelation,omitempty" binding:"omitempty,oneof=PARENT SPOUSE SIBLING FRIEND OTHER"`
	Status                   string                   `json:"status,omitempty" binding:"omitempty,oneof=ACTIVE INACTIVE DECEASED"`
}

// UpdatePatientDTO defines the fields allowed for updating a patient
type UpdatePatientDTO struct {
	PatientImage             *string                    `json:"patientImage,omitempty"`
	FullName                 *string                    `json:"fullName,omitempty"`
	Sex                      *Sex                       `json:"sex,omitempty" binding:"omitempty,oneof=MALE FEMALE"`
	DateOfBirth              *time.Time                 `json:"dateOfBirth,omitempty"`
	BloodType                *BloodType                 `json:"bloodType,omitempty" binding:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	Phone                    *string                    `json:"phone,omitempty" binding:"omitempty,len=10,numeric"`
	Email                    *string                    `json:"email,omitempty" binding:"omitempty,email"`
	City                     *string                    `json:"city,omitempty"`
	SubCity                  *string                    `json:"subcity,omitempty"`
	Wereda                   *string                    `json:"wereda,omitempty"`
	EmergencyContactName     *string                    `json:"emergencyContactName,omitempty"`
	EmergencyContactPhone    *string                    `json:"emergencyContactPhone,omitempty" binding:"omitempty,len=10,numeric"`
	EmergencyContactRelation *EmergencyContactRelation  `json:"emergencyContactRelation,omitempty" binding:"omitempty,oneof=PARENT SPOUSE SIBLING FRIEND OTHER"`
	Allergies                *[]patient_allergy.Allergy `json:"allergies,omitempty"`
	Status                   *string                    `json:"status,omitempty" binding:"omitempty,oneof=ACTIVE INACTIVE DECEASED"`
}
