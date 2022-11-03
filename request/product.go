package request

import "ecommerce/entities"

type ProductRequest struct {
	CategoryID  string             `json:"category_id" binding:"required"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	Price       int                `json:"price" binding:"required"`
	Variant     []entities.Variant `json:"variant"`
	CreatedBy   string
}

type VariantRequest struct {
	Sku        []string `form:"sku"`
	Colour     []string `form:"colour"`
	Size       []string `form:"size"`
	Ingredient []string `form:"ingredient"`
	Quantity   []string `form:"quantity"`
}
