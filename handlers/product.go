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

type productHandler struct {
	productService  services.ProductService
	categoryService services.CategoryService
	variantService  services.VariantService
	authService     middleware.ServiceAuth
	pagination      utils.Pagination
}

func NewProductHandler(authService middleware.ServiceAuth, productService services.ProductService, categoryService services.CategoryService, variantService services.VariantService, pagination utils.Pagination) *productHandler {
	return &productHandler{productService, categoryService, variantService, authService, pagination}
}

func (h *productHandler) Create(c *gin.Context) {
	var (
		input request.ProductRequest
	)
	bindErr := c.ShouldBindJSON(&input)
	fmt.Println("inputan", input.Variant)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, helpers.MessageBindRequest, bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	currentUser := c.MustGet("current_user").(entities.User)

	createProduct, err := h.productService.Create(currentUser.ID, input)
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, helpers.MessageFailedInsert, err)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	_, errVariant := h.variantService.Create(createProduct.ID, input)
	getCategory, _ := h.categoryService.FindByID(createProduct.CategoryID.String())
	getVariants, _ := h.variantService.FindVariantByProductID(createProduct.ID.String())

	if errVariant != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, helpers.MessageFailedInsert, errVariant)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, false, helpers.MessageSuccess, formatresponse.ConvProduct(createProduct, getCategory, getVariants))
	c.JSON(http.StatusCreated, result)
	return
}
func (h *productHandler) Product(c *gin.Context) {
	id := c.Param("id")
	findProduct, errFind := h.productService.FindProductByID(id)
	getCategory, _ := h.categoryService.FindByID(findProduct.CategoryID.String())
	if errFind != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, helpers.MessageFailed, errFind)
		c.JSON(http.StatusNotFound, result)
		return
	}
	if findProduct.ID.String() == "00000000-0000-0000-0000-000000000000" {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, helpers.MessageFailed, errFind)
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, helpers.StatusOK, helpers.MessageSuccess, formatresponse.NewConvProduct(findProduct, getCategory))
	c.JSON(http.StatusCreated, result)
	return
}
