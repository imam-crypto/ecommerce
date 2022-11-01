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
	"math/rand"
	"net/http"
)

type categoryHandler struct {
	authService     middleware.ServiceAuth
	categoryService services.CategoryService
	pagination      utils.Pagination
}

func NewCategoryHandlers(authService middleware.ServiceAuth, categoryService services.CategoryService, pagination utils.Pagination) *categoryHandler {
	return &categoryHandler{authService, categoryService, pagination}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZIMAM")

func code(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input request.CategoryRequestInsert
	bindErr := c.ShouldBind(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "check your input", bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	formImages, fileName, err := c.Request.FormFile("image")
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "failed", err)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	currentUser := c.MustGet("current_user").(entities.User)
	uploadUrl, publicID, err := h.categoryService.ImageUpload(request.CategoryImageRequest{Image: formImages}, fileName.Filename)
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "failed", err.Error())
		c.JSON(http.StatusBadRequest, result)
		return
	}
	fmt.Println("public ID nya dari Cloudinary [][][][]", publicID)
	input.PublicIDCloud = publicID

	createCategory, errCreate := h.categoryService.CreateCategory(currentUser.ID, input, uploadUrl)
	if errCreate != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, "failed", errCreate)
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusCreated, true, "success", formatresponse.ConvCategoryResponse(createCategory))
	c.JSON(http.StatusCreated, result)
	return
}

func (h *categoryHandler) Category(c *gin.Context) {
	id := c.Param("id")
	category, errGet := h.categoryService.FindByID(id)
	if errGet != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", errGet)
		c.JSON(http.StatusNotFound, result)
		return
	}
	if category.ID.String() == "00000000-0000-0000-0000-000000000000" {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", errGet)
		c.JSON(http.StatusNotFound, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "success", formatresponse.ConvCategoryResponse(category))
	c.JSON(http.StatusOK, result)
	return
}

func (h *categoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input request.CategoryRequestInsert
	bindErr := c.ShouldBind(&input)
	if bindErr != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "check your input", bindErr)
		c.JSON(http.StatusBadRequest, result)
		return
	}
	// get value from db
	oldCategory, errGet := h.categoryService.FindByID(id)
	if errGet != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", errGet)
		c.JSON(http.StatusNotFound, result)
		return
	}
	DelOldPublicID := oldCategory.PublicIDCloud
	// check input image when update
	formImages, fileName, err := c.Request.FormFile("image")
	if err != nil {
		updateCategory, errUpdate := h.categoryService.Update(id, input, oldCategory.UrlImage)
		if errUpdate != nil {
			result := helpers.ConvDefaultResponse(http.StatusNotModified, false, "failed", errUpdate)
			c.JSON(http.StatusNotModified, result)
			fmt.Println("gamabrnya", formImages)
			return
		}
		fmt.Println("gamabarnyya tidak di ganti")
		result := helpers.ConvDefaultResponse(http.StatusOK, true, "success", formatresponse.ConvCategoryResponse(updateCategory))
		c.JSON(http.StatusOK, result)
		return
	}
	// end check

	//	if user changes image
	//getPublicID := code(5)
	deleteOldImage, _ := helpers.FileDestroy(DelOldPublicID)

	secureUrl, publicID, err := h.categoryService.ImageUpload(request.CategoryImageRequest{Image: formImages}, fileName.Filename)
	input.PublicIDCloud = publicID
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusBadRequest, false, "failed", err.Error())
		c.JSON(http.StatusBadRequest, result)
		return
	}

	fmt.Println("gambarnya di hapus[][][][][]", deleteOldImage)
	updateCategory, errUpdate := h.categoryService.Update(id, input, secureUrl)
	if errUpdate != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotModified, false, "failed", errUpdate)
		c.JSON(http.StatusNotModified, result)
		fmt.Println("gamabrnya", formImages)
		return
	}
	fmt.Println("gambarnya di ganti")
	result := helpers.ConvDefaultResponse(http.StatusOK, true, "success", formatresponse.ConvCategoryResponse(updateCategory))
	c.JSON(http.StatusOK, result)
	return
}
func (h *categoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	currentUser := c.MustGet("current_user").(entities.User)

	// del image from cloudinary
	getCategory, errGet := h.categoryService.FindByID(id)
	if errGet != nil {
		result := helpers.ConvDefaultResponse(http.StatusNotFound, false, "failed", errGet)
		c.JSON(http.StatusNotFound, result)
		return
	}
	urlImage := getCategory.UrlImage

	deleteOldImage, _ := helpers.FileDestroy(urlImage)
	fmt.Println("gambarnya di hapus[][][][][]", deleteOldImage)
	fmt.Println("result url cloud", urlImage)
	_, err := h.categoryService.Delete(id, currentUser.ID)
	if err != nil {
		result := helpers.ConvDefaultResponse(http.StatusUnprocessableEntity, false, "failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	result := helpers.ConvDefaultResponse(http.StatusGone, true, "success", deleteOldImage)
	c.JSON(http.StatusGone, result)
}
func (h *categoryHandler) Categories(c *gin.Context) {
	var searchFilter string
	queries := c.Request.URL.Query()
	for queryKey, queryValue := range queries {
		if queryKey == "search" {
			searchFilter = queryValue[len(queryValue)-1]
		}
	}
	pagination := h.pagination.GetPagination(c)
	category, pagination := h.categoryService.GetAllCategory(searchFilter, pagination)

	var categoryRes []formatresponse.CategoryResponse
	for _, pi := range category {
		res := formatresponse.ConvCategoryResponse(pi)
		categoryRes = append(categoryRes, res)
	}

	result := helpers.ConvResponsePaginate(http.StatusOK, true, "Success", categoryRes, pagination)
	c.JSON(http.StatusOK, result)
}
