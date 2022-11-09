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

func RoleRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	//userRepo := repositories.NewUserRepository(db)
	//userService := services.NewUserService(userRepo)

	roleRepo := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepo)

	RoleHandler := handlers.NewRoleHandlers(authService, roleService, *pagination)
	router.POST("/role/create", RoleHandler.Create)

}
