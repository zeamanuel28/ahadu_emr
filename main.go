package main

import (
	"log"
	"saas-api/core"
	"saas-api/modules/allergy"
	"saas-api/modules/auth"
	"saas-api/modules/branch"
	"saas-api/modules/department"
	"saas-api/modules/employee"
	"saas-api/modules/patient"
	"saas-api/modules/position"
	"saas-api/modules/user"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Application
	app := core.NewApp()

	// Register Modules
	app.RegisterModule(auth.NewModule())
	app.RegisterModule(user.NewModule())
	app.RegisterModule(allergy.NewModule())
	app.RegisterModule(branch.NewModule())
	app.RegisterModule(department.NewModule())
	app.RegisterModule(position.NewModule())
	app.RegisterModule(employee.NewModule())
	app.RegisterModule(patient.NewModule())

	// Initialize and Run
	app.Init()
	app.Run()
}
