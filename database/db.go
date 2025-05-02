package database

import (
	"clinic-management/backend/config"
	"clinic-management/backend/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	// "github.com/denisenkom/go-mssqldb" // Import MSSQL driver
	"gorm.io/gorm"
)

var DB *gorm.DB

// Initialize the database connection using GORM
func Initialize() {
	var err error

	config.ReadConfig()

	cfg := config.GetConfig()

	// Build the connection string from the config file
	connString := fmt.Sprintf(
		cfg.DbServer.ConnectionString,
	)

	// Connect to the SQL Server database using GORM
	DB, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Check the connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	// Ping the database
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	DB.AutoMigrate(&models.CMSUser{})
	DB.AutoMigrate(&models.CMSPatients{})
	DB.AutoMigrate(&models.CMSRolesMaster{})
	DB.AutoMigrate(&models.CMSUserRoles{})
	DB.AutoMigrate(&models.CMSEmployeeMaster{})
	DB.AutoMigrate(&models.CMSDepartments{})
	DB.AutoMigrate(&models.CMSAutoNumber{})
	DB.AutoMigrate(&models.CMSEmployeeLeaveMaster{})


	log.Println("Connected to SQL Server using GORM")
	fmt.Println("Connected to SQL Server using GORM")
}
