package handlers

import (
	"ecommerce/dtos"
	"ecommerce/entities"
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
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, helpers.StatusFailed, helpers.MessageFailedInsert, errVariant)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, helpers.StatusOK, helpers.MessageSuccess, dtos.ConvProduct(createProduct, getCategory, getVariants))
	c.JSON(http.StatusCreated, result)
	return
}
func (h *productHandler) Product(c *gin.Context) {
	id := c.Param("id")
	findProduct, errFind := h.productService.FindProductByID(id)
	getCategory, _ := h.categoryService.FindByID(findProduct.CategoryID.String())
	if errFind != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, helpers.StatusFailed, helpers.MessageFailed, errFind)
		c.JSON(http.StatusNotFound, result)
		return
	}
	if findProduct.ID.String() == "00000000-0000-0000-0000-000000000000" {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, helpers.StatusFailed, helpers.MessageFailed, errFind)
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, helpers.StatusOK, helpers.MessageSuccess, dtos.NewConvProduct(findProduct, getCategory))
	c.JSON(http.StatusCreated, result)
	return
}

func (h *productHandler) GetVariants(c *gin.Context) {
	id := c.Param("id")
	variant, errGet := h.productService.FindProductByID(id)
	if errGet != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, helpers.StatusFailed, helpers.MessageFailed, errGet)
		c.JSON(http.StatusNotFound, result)
		return
	}
	if variant.ID.String() == "00000000-0000-0000-0000-000000000000" {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, helpers.StatusFailed, helpers.MessageNotFound, helpers.MessageNotFound)
		c.JSON(http.StatusNotFound, result)
		return
	}
	getCategory, _ := h.categoryService.FindByID(variant.CategoryID.String())
	getVariants, _ := h.variantService.FindVariantByProductID(variant.ID.String())
	result := helpers.ConvDefaultResponse(http.StatusOK, helpers.StatusOK, helpers.MessageSuccess, dtos.ConvProduct(variant, getCategory, getVariants))
	c.JSON(http.StatusOK, result)
	return
}

func (h *productHandler) Update(c *gin.Context) {
	var (
		input request.ProductRequestUpdate
	)
	id := c.Param("id")
	bindErr := c.ShouldBindJSON(&input)
	//fmt.Println("inputan", input.Variant)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, helpers.StatusFailed, helpers.MessageBindRequest, bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	updateProduct, errUpdate := h.productService.Update(id, input)
	if errUpdate != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotModified, helpers.StatusFailed, helpers.MessageFailed, errUpdate)
		c.JSON(http.StatusNotModified, result)
		return
	}

	getCategory, _ := h.categoryService.FindByID(updateProduct.CategoryID.String())
	getVariants, _ := h.variantService.FindVariantByProductID(updateProduct.ID.String())
	result := helpers.ConvDefaultResponse(http.StatusOK, helpers.StatusOK, helpers.MessageSuccess, dtos.ConvProduct(updateProduct, getCategory, getVariants))
	c.JSON(http.StatusOK, result)
	return
}

func (h *productHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	currentUser := c.MustGet("current_user").(entities.User)

	del, err := h.productService.Delete(id, currentUser.ID)
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, "failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	varDel, _ := h.variantService.DeleteByIdProduct(id, currentUser.ID)
	fmt.Println("variant delete", varDel)
	result := helpers.ConvDefaultResponse(http.StatusGone, true, helpers.MessageSuccess, del)
	c.JSON(http.StatusGone, result)
}

func (h *productHandler) Products(c *gin.Context) {
	var searchFilter string
	queries := c.Request.URL.Query()
	for queryKey, queryValue := range queries {
		if queryKey == "search" {
			searchFilter = queryValue[len(queryValue)-1]
		}
	}
	pagination := h.pagination.GetPagination(c)
	product, pagination := h.productService.GetAllProducts(searchFilter, pagination)

	var products []dtos.ProductResponse
	for _, productResp := range product {
		res := dtos.ConvProductresponsePaginate(productResp)
		products = append(products, res)
	}

	result := helpers.ConvResponsePaginate(http.StatusOK, true, helpers.MessageSuccess, products, pagination)
	c.JSON(http.StatusOK, result)
}
