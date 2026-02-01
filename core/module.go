package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Module defines the interface for a project module
type Module interface {
	// Init initializes the module with the database connection
	Init(db *gorm.DB)

	// RegisterRoutes registers the module's routes to the Gin group
	RegisterRoutes(parentGroup *gin.RouterGroup)

	// GetModels returns the GORM models managed by this module for auto-migration
	GetModels() []interface{}
}
