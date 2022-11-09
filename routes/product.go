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

func ProductRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	variantRepo := repositories.NewVariantRepository(db)
	variantService := services.NewVariantService(variantRepo)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(authService, productService, categoryService, variantService, *pagination)
	router.POST("/product/create", middleware.AuthAdmin(authService, userService), productHandler.Create)
	router.GET("/product/:id", middleware.AuthAdmin(authService, userService), productHandler.Product)
	router.PUT("/product/update/:id", middleware.AuthAdmin(authService, userService), productHandler.Update)
}
