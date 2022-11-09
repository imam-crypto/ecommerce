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

type menuAccessHandler struct {
	authService       middleware.ServiceAuth
	menuAccessService services.MenuAccessService
	pagination        utils.Pagination
}

func NewMenuAccessHandlers(authService middleware.ServiceAuth, menuAccessService services.MenuAccessService, pagination utils.Pagination) *menuAccessHandler {
	return &menuAccessHandler{authService, menuAccessService, pagination}
}

func (h *menuAccessHandler) Create(c *gin.Context) {
	var input request.MenuAccessRequest
	bindErr := c.ShouldBind(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, helpers.MessageBindRequest, bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	currentUser := c.MustGet("current_user").(entities.User)

	createAccess, errAccess := h.menuAccessService.CreateMenuAccess(currentUser.ID, input)
	if errAccess != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, helpers.MessageFailed, errAccess)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, true, helpers.MessageSuccess, createAccess)
	c.JSON(http.StatusCreated, result)
	return

}
