package main

import (
	"log"
	"saas-api/core"
	"saas-api/modules/allergy_name"
	"saas-api/modules/allergy_reaction"
	"saas-api/modules/auth"
	"saas-api/modules/branch"
	"saas-api/modules/department"
	"saas-api/modules/diagnosis"
	"saas-api/modules/diagnosis_code"
	"saas-api/modules/disposition"
	"saas-api/modules/employee"
	"saas-api/modules/observation"
	"saas-api/modules/patient"
	"saas-api/modules/patient_allergy"
	"saas-api/modules/position"
	"saas-api/modules/problem"
	"saas-api/modules/user"
	"saas-api/modules/visit"
	"saas-api/modules/vital_record"
	"saas-api/modules/vital_type"

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
	app.RegisterModule(observation.NewModule())
	app.RegisterModule(vital_record.NewModule())
	app.RegisterModule(problem.NewModule())
	app.RegisterModule(vital_type.NewModule())
	app.RegisterModule(diagnosis_code.NewModule())
	app.RegisterModule(diagnosis.NewModule())
	app.RegisterModule(disposition.NewModule())

	// Initialize and Run
	app.Init()
	app.Run()
}
