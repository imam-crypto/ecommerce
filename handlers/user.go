package handlers

import (
	"ecommerce/formatresponse"
	"ecommerce/helpers"
	"ecommerce/middleware"
	"ecommerce/request"
	"ecommerce/services"
	"ecommerce/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
func (h *userHandler) Regiter(c *gin.Context) {

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

	fmt.Println("ID string nya", loogedinUser.ID.String())

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
