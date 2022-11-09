package dtos

import "ecommerce/entities"

type ProductResponse struct {
	ID          string            `json:"id"`
	Category    CategoryResponse  `json:"category"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Variant     []VariantResponse `json:"variant"`
}

type VariantResponse struct {
	ID         string `json:"id"`
	Colour     string `json:"colour"`
	Size       string `json:"size"`
	Quantity   int    `json:"quantity"`
	Ingredient string `json:"ingredient"`
}

func ConvProduct(product entities.Product, ct entities.Category, variants []entities.Variant) ProductResponse {
	formatProduct := ProductResponse{
		ID:          product.ID.String(),
		Category:    ConvCategoryResponse(ct),
		Title:       product.Title,
		Description: product.Description,
	}
	for _, variantObj := range variants {
		variantRes := ConvVariantResponse(variantObj)
		formatProduct.Variant = append(formatProduct.Variant, variantRes)
	}
	return formatProduct
}
func NewConvProduct(product entities.Product, ct entities.Category) ProductResponse {
	formatProduct := ProductResponse{
		ID:          product.ID.String(),
		Category:    ConvCategoryResponse(ct),
		Title:       product.Title,
		Description: product.Description,
	}
	for _, variantObj := range product.Variant {
		variantRes := ConvVariantResponse(variantObj)
		formatProduct.Variant = append(formatProduct.Variant, variantRes)
	}
	return formatProduct
}

func ConvVariantResponse(variant entities.Variant) VariantResponse {
	formatVariant := VariantResponse{
		ID:         variant.ID.String(),
		Colour:     variant.Colour,
		Size:       variant.Size,
		Ingredient: variant.Ingredient,
		Quantity:   variant.Quantity,
	}
	return formatVariant
}
