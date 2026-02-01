package core

import (
	"log"
	"net/http"
	"os"
	"saas-api/shared/database"
	"saas-api/shared/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type App struct {
	Router  *gin.Engine
	DB      *gorm.DB
	Modules []Module
}

func NewApp() *App {
	return &App{
		Router:  gin.Default(),
		Modules: []Module{},
	}
}

func (a *App) RegisterModule(m Module) {
	a.Modules = append(a.Modules, m)
}

func (a *App) Init() {
	// Initialize Database
	database.InitDB()
	a.DB = database.DB

	// Collect models for migration
	var allModels []interface{}
	for _, m := range a.Modules {
		m.Init(a.DB)
		allModels = append(allModels, m.GetModels()...)
	}

	// Auto-migrate
	// Auto-migrate
	if len(allModels) > 0 {
		err := a.DB.AutoMigrate(allModels...)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	// Swagger setup - fully dynamic
	a.Router.GET("/swagger-api/doc.json", func(c *gin.Context) {
		// Register Response schema explicitly to ensure it exists for references
		utils.Registry.RegisterSchema("utils.Response", utils.Response{})

		spec, err := utils.Registry.GenerateFullSpec(utils.SwaggerConfig{
			Title:       "SaaS API Implementation in Go",
			Description: "This is a migrated SaaS API using Gin and GORM.",
			Version:     "1.0",
			Host:        "localhost:8080",
			BasePath:    "/api/v1",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Header("Content-Type", "application/json")
		c.String(http.StatusOK, spec)
	})
	a.Router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger-api/doc.json")))

	// Global Middleware
	a.Router.Use(gin.Recovery())
	a.Router.Use(gin.Logger())

	// API Group
	api := a.Router.Group("/api/v1")

	// Register module routes
	for _, m := range a.Modules {
		m.RegisterRoutes(api)
	}
}

func (a *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.Router.Run(":" + port)
}
