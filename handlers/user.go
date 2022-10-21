package handlers

import (
	"ecommerce/services"
	"ecommerce/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService services.UserServices
	pagination  utils.Pagination
}

func NewUserHandlers(userService services.UserServices, pagination utils.Pagination) *userHandler {
	return &userHandler{userService, pagination}
}

func (h *userHandler) GetUser(c *gin.Context) {

	id := c.Param("id")
	fmt.Print(id)
}
