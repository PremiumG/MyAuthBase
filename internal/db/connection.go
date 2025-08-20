package db

import (
	"AuthBase/internal/models"
	"AuthBase/internal/utils"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Initialize() error {

	var err error
	// Open SQLite database with GORM
	DB, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s/db.db?_foreign_keys=on", utils.AppConfig.DBPath)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

func Migrate() error {
	// Auto-migrate all models
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migration completed")
	return nil
}

func SeedAdmin(email string) error {
	admin := models.User{
		Email:   email,
		IsAdmin: true,
	}

	// Create only if doesn't exist
	result := DB.Where("email = ?", email).FirstOrCreate(&admin)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Admin user seeded: %s", email)
	return nil
}

func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
