package main

import (
	"log"
	"saas-api/core"
	"saas-api/modules/allergy_name"
	"saas-api/modules/allergy_reaction"
	"saas-api/modules/auth"
	"saas-api/modules/branch"
	"saas-api/modules/department"
	"saas-api/modules/employee"
	"saas-api/modules/patient"
	"saas-api/modules/patient_allergy"
	"saas-api/modules/position"
	"saas-api/modules/user"
	"saas-api/modules/visit"

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
	app.RegisterModule(patient.NewModule())
	app.RegisterModule(allergy_name.NewModule())
	app.RegisterModule(allergy_reaction.NewModule())
	app.RegisterModule(patient_allergy.NewModule())
	app.RegisterModule(branch.NewModule())
	app.RegisterModule(department.NewModule())
	app.RegisterModule(position.NewModule())
	app.RegisterModule(employee.NewModule())
	app.RegisterModule(visit.NewModule())

	// Initialize and Run
	app.Init()
	app.Run()
}
