package main

import (
	"ecommerce/routes"
	"ecommerce/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config, db, err := configInitalization()
	if err != nil {
		log.Fatal("Config initialization failed.")
	}

	// Migration: Set GORM_MIGRATE equal to true in app.env file if you want to migrate all models
	if config.GormMigrate == "true" {
		utils.MigrateModels(db)
	}

	// Gin init
	router := utils.SetupRouter(config)
	// Setup API groups or versions
	v1 := router.Group("/v1")

	// Routes V1
	routes.UserRoute(&config, db, v1)
}

func configInitalization() (utils.Config, *gorm.DB, error) {
	var (
		config utils.Config
		db     *gorm.DB
		err    error
	)

	// Load configuration file
	config, err = utils.LoadConfig(".", false)

	if err != nil {
		return config, db, err
	}
	// Connect to database
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort, config.DbTz,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return config, db, err
	}
	log.Println("[DATABASE] Database connection success.")
	return config, db, err
}
