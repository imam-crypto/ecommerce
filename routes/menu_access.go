package routes

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"ecommerce/repositories"
	"ecommerce/services"
	"ecommerce/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuAccessRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	menuAccessRepository := repositories.NewMenuAccessRepository(db)
	menuAccessService := services.NewMenuAccessService(menuAccessRepository)
	menuAccessHandler := handlers.NewMenuAccessHandlers(authService, menuAccessService, *pagination)
	router.POST("/menu-access/create", middleware.AuthAdmin(authService, userService), menuAccessHandler.Create)
}
