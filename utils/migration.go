package utils

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// Register model/entity here to migrate.
	// Set GORM_MIGRATE equal to true in app.env file to run migration on server start

	db.AutoMigrate(

		&entities.Category{},

		&entities.User{},
	)

}
