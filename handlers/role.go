package handlers

import (
	"ecommerce/helpers"
	"ecommerce/middleware"
	"ecommerce/request"
	"ecommerce/services"
	"ecommerce/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type RoleHandler struct {
	authService middleware.ServiceAuth
	roleService services.RoleService
	pagination  utils.Pagination
}

func NewRoleHandlers(authService middleware.ServiceAuth, roleService services.RoleService, pagination utils.Pagination) *RoleHandler {
	return &RoleHandler{authService, roleService, pagination}
}

func (h *RoleHandler) Create(c *gin.Context) {
	var input request.RoleRequest
	bindErr := c.ShouldBindJSON(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, helpers.StatusFailed, helpers.MessageFailed, bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	//currentUser := c.MustGet("current_user").(entities.User)
	createRole, errCreate := h.roleService.Create(uuid.FromStringOrNil("1a71d76e-090b-41d8-86b5-a36e9f7cb01c"), input)
	if errCreate != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, helpers.StatusFailed, helpers.MessageFailed, errCreate)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, helpers.StatusOK, helpers.MessageSuccess, createRole)
	c.JSON(http.StatusCreated, result)
	return
}
