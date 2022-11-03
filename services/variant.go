package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type VariantService interface {
	Create(productID uuid.UUID, input request.ProductRequest) ([]entities.Variant, error)
	FindVariantByProductID(productID string) ([]entities.Variant, error)
}
type varianService struct {
	variantRepositories repositories.VariantRepository
}

func NewVariantService(variantRepositories repositories.VariantRepository) *varianService {
	return &varianService{variantRepositories}
}
func (s *varianService) Create(productID uuid.UUID, input request.ProductRequest) ([]entities.Variant, error) {
	var variants = []entities.Variant{}

	for _, varianResp := range input.Variant {
		create := entities.Variant{
			ProductID:  productID,
			Sku:        varianResp.Sku,
			Colour:     varianResp.Colour,
			Size:       varianResp.Size,
			Ingredient: varianResp.Ingredient,
			Quantity:   varianResp.Quantity,
		}
		variants = append(variants, create)
		createVariant, errCreate := s.variantRepositories.Create(create)
		if errCreate != nil {
			return variants, errCreate
		}
		fmt.Println("create varian di service variant", createVariant)
	}
	return variants, nil
}
func (s *varianService) FindVariantByProductID(productID string) ([]entities.Variant, error) {

	find, err := s.variantRepositories.FindVariantByProductID(productID)
	if err != nil {
		return find, err
	}
	return find, nil
}
