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

func MenuRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	menuRepo := repositories.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)

	menuHandler := handlers.NewMenuHandlers(authService, menuService, *pagination)
	router.GET("/menus", menuHandler.Create)
	router.POST("/menu/create", middleware.AuthAdmin(authService, userService), menuHandler.Create)
}
