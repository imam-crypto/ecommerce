package utils

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(config Config) *gin.Engine {
	// Start Gin router.
	// GIN_MODE can be switched in app.env file:
	// Set GIN_MODE to "debug" if you are in development environment.
	// Set GIN_MODE to "release" if you are in production environment.
	// DO NOT USE "debug" IN PRODUCTION!
	if config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFS("/file", http.Dir("public"))

	return router
}
