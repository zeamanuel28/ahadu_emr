package user

import (
	"saas-api/core"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
)

type User struct {
	core.BaseModel
	Username     string     `gorm:"not null" json:"username"`
	Email        string     `gorm:"uniqueIndex;not null" json:"email"`
	FullName     string     `json:"fullName"`
	Password     string     `gorm:"not null" json:"-"`
	TempPassword string     `json:"tempPassword,omitempty"`
	Status       UserStatus `gorm:"type:string;default:ACTIVE" json:"status"`
	Roles        []Role     `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles,omitempty"`
}

type Role struct {
	core.BaseModel
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
}
