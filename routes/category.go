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

func CategoryRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandlers(authService, categoryService, *pagination)
	router.GET("/categories", categoryHandler.Categories)
	router.POST("/category/create", middleware.AuthAdmin(authService, userService), categoryHandler.CreateCategory)
	router.GET("/category/:id", middleware.AuthAdmin(authService, userService), categoryHandler.Category)
	router.PUT("/category/update/:id", middleware.AuthAdmin(authService, userService), categoryHandler.Update)
	router.DELETE("/category/delete/:id", middleware.AuthAdmin(authService, userService), categoryHandler.Delete)
}
