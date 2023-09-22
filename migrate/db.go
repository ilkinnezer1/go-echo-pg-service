package migrate

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/models"
	"os"
)

func ConnectDatabase() (*gorm.DB, error) {
	err := godotenv.Load()
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate the schema for Blog
	if err := db.AutoMigrate(models.Blog{}); err != nil {
		return nil, err
	}
	// AutoMigrate the schema for Partners
	if err := db.AutoMigrate(models.Partners{}); err != nil {
		return nil, err
	}
	// AutoMigrate the schema for Projects
	if err := db.AutoMigrate(models.Projects{}); err != nil {
		return nil, err
	}
	// AutoMigrate User Credentials
	if err := db.AutoMigrate(models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
