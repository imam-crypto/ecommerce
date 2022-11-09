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
type ProductRequestUpdate struct {
	CategoryID  string                 `json:"category_id" binding:"required"`
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description" binding:"required"`
	Price       int                    `json:"price" binding:"required"`
	Variant     []VariantRequestUpdate `json:"variant"`
	CreatedBy   string
}
type VariantRequestUpdate struct {
	ID         string `json:"id"`
	ProductID  string `json:"product_id"`
	Sku        string `json:"sku"`
	Colour     string `json:"colour"`
	Size       string `json:"size"`
	Ingredient string `json:"ingredient"`
	Quantity   int    `json:"quantity"`
}

type VariantRequest struct {
	Sku        []string `form:"sku"`
	Colour     []string `form:"colour"`
	Size       []string `form:"size"`
	Ingredient []string `form:"ingredient"`
	Quantity   []string `form:"quantity"`
}
