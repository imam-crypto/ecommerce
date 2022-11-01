package handlers

import (
	"ecommerce/entities"
	"ecommerce/formatresponse"
	"ecommerce/helpers"
	"ecommerce/middleware"
	"ecommerce/request"
	"ecommerce/services"
	"ecommerce/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	authService middleware.ServiceAuth
	userService services.UserServices
	pagination  utils.Pagination
}

func NewUserHandlers(authService middleware.ServiceAuth, userService services.UserServices, pagination utils.Pagination) *userHandler {
	return &userHandler{authService, userService, pagination}
}

func (h *userHandler) GetUser(c *gin.Context) {

	id := c.Param("id")
	fmt.Print(id)
}

func (h *userHandler) Register(c *gin.Context) {
	var input request.RegisterUserInput
	bindErr := c.ShouldBindJSON(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "email has been used", "failed Register")
		c.JSON(http.StatusBadRequest, result)
		return
	}
	check, _ := h.userService.CheckEmail(input.Email)
	if check.Email == input.Email {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "email has been used", "failed Register")
		c.JSON(http.StatusBadRequest, result)
		return
	}
	newUser, err := h.userService.Register(input)
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, "failed", "failed Register")
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", formatresponse.ConvResponseUser(newUser))
	c.JSON(http.StatusOK, result)
}

func (h *userHandler) Login(c *gin.Context) {
	var input request.LoginUserInput
	bindErr := c.ShouldBindJSON(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "check your input", "failed login")
		c.JSON(http.StatusBadRequest, result)
		return
	}
	loogedinUser, erCek := h.userService.Login(input)
	if erCek != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", erCek)
		c.JSON(http.StatusNotFound, result)
		return
	}
	//fmt.Println("ID string nya", loogedinUser.ID.String())
	token, exp, errToken := h.authService.GenerateToken(loogedinUser.ID.String())
	// exp := token.ExpiresAt
	if errToken != nil {
		result := helpers.ConvDefaultResponse(http.StatusInternalServerError, false, "failed", errToken)
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	start := exp.Format("02-01-2006 15:04:05")
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", formatresponse.ConvResponseUserLogin(loogedinUser, token, start))
	c.JSON(http.StatusOK, result)
}
func (h *userHandler) Profile(c *gin.Context) {
	currentUser := c.MustGet("current_user").(entities.User)
	// role := currentUser.Role
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", formatresponse.ConvResponseUser(currentUser))
	c.JSON(http.StatusOK, result)
}

func (h *userHandler) Update(c *gin.Context) {
	var input request.UpdateUserInput
	bindErr := c.ShouldBindJSON(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "check your input", "failed update")
		c.JSON(http.StatusBadRequest, result)
		return
	}
	currentUser := c.MustGet("current_user").(entities.User)
	fmt.Println("id di update", currentUser.ID.String())
	updateUser, errUpdate := h.userService.UpdateUser(currentUser.ID.String(), input)
	if errUpdate != nil {
		result := helpers.ConvDefaultResponse(http.StatusInternalServerError, false, "failed update", errUpdate)
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", formatresponse.ConvResponseUser(updateUser))
	c.JSON(http.StatusOK, result)
	return
}

func (h *userHandler) GetUsers(c *gin.Context) {
	res, err := h.userService.GetUsers()
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", err)
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", res)
	c.JSON(http.StatusOK, result)
}
func (h *userHandler) GetAllUsers(c *gin.Context) {
	var searchFilter string
	queries := c.Request.URL.Query()
	for queryKey, queryValue := range queries {
		if queryKey == "search" {
			searchFilter = queryValue[len(queryValue)-1]
		}
	}
	pagination := h.pagination.GetPagination(c)
	users, pagination := h.userService.FindUserAllPaginate(searchFilter, pagination)

	var piRes []formatresponse.ResponseUser
	for _, pi := range users {
		res := formatresponse.PaginateUserResponse(pi)
		piRes = append(piRes, res)
	}
	result := helpers.ConvResponsePaginate(http.StatusOK, true, "Success", piRes, pagination)
	c.JSON(http.StatusOK, result)
}
func (h *userHandler) UpdateRole(c *gin.Context) {
	var input request.UpdateUserRole
	bindErr := c.ShouldBindJSON(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "check your input", "failed update")
		c.JSON(http.StatusBadRequest, result)
		return
	}
	currentUser := c.MustGet("current_user").(entities.User)
	fmt.Println("id di update", currentUser.ID.String())
	updateUser, errUpdate := h.userService.UpdateUserRole(currentUser.ID.String(), input)
	if errUpdate != nil {
		result := helpers.ConvDefaultResponse(http.StatusInternalServerError, false, "failed update", errUpdate)
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "Success", formatresponse.ConvResponseUser(updateUser))
	c.JSON(http.StatusOK, result)
	return
}
