package routes

import (
	"ecommerce/handlers"
	"ecommerce/repositories"
	"ecommerce/services"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(config *utils.Config, db *gorm.DB, router *gin.RouterGroup) {

	// bookRepo := repositories.NewBookRepositories(db)
	// bookService := services.NewBookService(bookRepo)
	// bookHandler := handlers.NewBookHandler(bookService)
	// authService := middleware.NewService()
	pagination := utils.NewPagination()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userhandler := handlers.NewUserHandlers(userService, *pagination)

	// router.POST("/user/register", userhandler.Register)
	// router.POST("/user/login-member", userhandler.Login)
	// router.POST("/user/login-admin", userhandler.Login)
	// router.GET("/user/forgot-password", userhandler.ForgotPassword)
	// router.POST("/user/reset-password", userhandler.ResetPassword)
	// router.GET("/user/report", userhandler.ReportUser)
	// router.GET("/user/profile", middleware.AuthUser(authService, userService), userhandler.Profile)
	router.GET("/user-get/:id", userhandler.GetUser)
	// router.GET("/user/admin-profile", middleware.AuthAdmin(authService, userService), userhandler.Profile)
	// router.PUT("/user/update", middleware.AuthUser(authService, userService), userhandler.Update)
	// router.PUT("/user/admin-update", middleware.AuthAdmin(authService, userService), userhandler.Update)
	// router.PUT("/user/update-role/:id", middleware.AuthAdmin(authService, userService), userhandler.UpdateRole)
	// router.POST("/user/logout", middleware.AuthUser(authService, userService), userhandler.Logout)
	// router.DELETE("/user/delete/:id", middleware.AuthAdmin(authService, userService), userhandler.Delete)
	// router.GET("/users", middleware.AuthAdmin(authService, userService), userhandler.GetAllUsers)

}
