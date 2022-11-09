package utils

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	// Register model/entity here to migrate.
	// Set GORM_MIGRATE equal to true in app.env file to run migration on server start

	db.AutoMigrate(
		&entities.Role{},
		&entities.User{},
		&entities.Category{},
		&entities.Menu{},
		&entities.MenuAccess{},
		&entities.Product{},
		&entities.Variant{},
		//&entities.ProductImage{},
	)

}
