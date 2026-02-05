package main

import (
	"fmt"
	"log"

	"saas-api/modules/branch"
	"saas-api/modules/department"
	"saas-api/modules/employee"
	"saas-api/modules/patient"
	"saas-api/modules/position"
	"saas-api/modules/user"
	"saas-api/shared/database"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize Database
	database.InitDB()

	// Auto-migrate models
	err := database.DB.AutoMigrate(
		&user.Role{},
		&user.User{},
		&branch.Branch{},
		&department.Department{},
		&position.Position{},
		&employee.Employee{},
		&patient.Patient{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	seedData()
}

func seedData() {
	fmt.Println("Seeding initial data...")

	// 1. Roles
	adminRole := user.Role{Name: "ADMIN", Description: "System Administrator"}
	userRole := user.Role{Name: "USER", Description: "Standard User"}

	database.DB.FirstOrCreate(&adminRole, user.Role{Name: "ADMIN"})
	database.DB.FirstOrCreate(&userRole, user.Role{Name: "USER"})

	// 2. Admin User
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	adminUser := user.User{
		Username: "admin",
		Email:    "admin@example.com",
		FullName: "System Administrator",
		Password: string(hashedPassword),
		Status:   user.UserStatusActive,
		Roles:    []user.Role{adminRole},
	}

	var existingUser user.User
	if err := database.DB.Where("email = ?", adminUser.Email).First(&existingUser).Error; err != nil {
		database.DB.Create(&adminUser)
		fmt.Println("Admin user created: admin@example.com / admin123")
	} else {
		fmt.Println("Admin user already exists")
	}

	// 3. Default Organizational Data
	defaultBranch := branch.Branch{Name: "Main Branch", Code: "MB001"}
	database.DB.Where(branch.Branch{Code: "MB001"}).FirstOrCreate(&defaultBranch)

	defaultDept := department.Department{Name: "IT Department", Code: "IT001"}
	database.DB.Where(department.Department{Code: "IT001"}).FirstOrCreate(&defaultDept)

	defaultPos := position.Position{Title: "Senior Developer", Code: "SD001", IsManagerial: true}
	database.DB.Where(position.Position{Code: "SD001"}).FirstOrCreate(&defaultPos)

	fmt.Println("Seed completed successfully!")
}
