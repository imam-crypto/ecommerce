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

func UserRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userhandler := handlers.NewUserHandlers(authService, userService, *pagination)

	router.POST("/user/register", userhandler.Register)
	router.POST("/user/login-member", userhandler.Login)
	// router.POST("/user/login-admin", userhandler.Login)
	// router.GET("/user/forgot-password", userhandler.ForgotPassword)
	// router.POST("/user/reset-password", userhandler.ResetPassword)
	router.GET("/user/users-cache", userhandler.GetUsers)
	router.GET("/user/users", middleware.AuthAdmin(authService, userService), userhandler.GetAllUsers)
	// router.GET("/user/profile", middleware.AuthUser(authService, userService), userhandler.Profile)
	router.GET("/user-get/:id", userhandler.GetUser)
	router.GET("/user/admin-profile", middleware.AuthAdmin(authService, userService), userhandler.Profile)
	router.PUT("/user/update", middleware.AuthAdmin(authService, userService), userhandler.Update)
	router.PUT("/user/admin-update", middleware.AuthAdmin(authService, userService), userhandler.Update)
	//router.PUT("/user/update-role/:id", middleware.AuthAdmin(authService, userService), userhandler.UpdateRole)
	// router.POST("/user/logout", middleware.AuthUser(authService, userService), userhandler.Logout)
	// router.DELETE("/user/delete/:id", middleware.AuthAdmin(authService, userService), userhandler.Delete)
	// router.GET("/users", middleware.AuthAdmin(authService, userService), userhandler.GetAllUsers)

}
