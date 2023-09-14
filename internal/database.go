package internal

import (
	"github.com/r35krag0th/datasorcerer/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Slot{},
		&models.Realm{},
		&models.CharacterClass{},
		&models.Build{},
		&models.Character{},
		&models.Weight{},
	)
}
