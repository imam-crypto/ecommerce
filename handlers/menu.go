package handlers

import (
	"ecommerce/entities"
	"ecommerce/helpers"
	"ecommerce/middleware"
	"ecommerce/request"
	"ecommerce/services"
	"ecommerce/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type menuHandler struct {
	authService middleware.ServiceAuth
	menuService services.MenuService
	pagination  utils.Pagination
}

func NewMenuHandlers(authService middleware.ServiceAuth, menuService services.MenuService, pagination utils.Pagination) *menuHandler {
	return &menuHandler{authService, menuService, pagination}
}

func (h *menuHandler) Create(c *gin.Context) {
	var input request.MenuRequest
	bindErr := c.ShouldBind(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, helpers.StatusFailed, helpers.MessageBindRequest, bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	currentUser := c.MustGet("current_user").(entities.User)
	_, errCreate := h.menuService.Create(currentUser.ID, input)
	if errCreate != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, helpers.StatusFailed, helpers.MessageFailed, errCreate)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, helpers.StatusOK, helpers.MessageSuccess, helpers.StatusOK)
	c.JSON(http.StatusCreated, result)
	return
}
