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
	Update(productID string, input request.ProductRequestUpdate) ([]request.VariantRequestUpdate, error)
	DeleteByIdProduct(idProduct string, loggedUser uuid.UUID) (bool, error)
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
		_, errCreate := s.variantRepositories.Create(create)
		if errCreate != nil {
			return variants, errCreate
		}
		//fmt.Println("create varian di service variant", createVariant)
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
func (s *varianService) Update(productID string, input request.ProductRequestUpdate) ([]request.VariantRequestUpdate, error) {

	// check existing data || delete
	checkExitingDatas := s.checkDataExisting(productID, input)
	fmt.Println("existing datas", checkExitingDatas)
	for _, variant := range input.Variant {
		variantData, _ := s.variantRepositories.FindVariantByID(variant.ID)
		if variantData.ID != uuid.Nil {
			variantData.Sku = variant.Sku
			variantData.Colour = variant.Colour
			variantData.Size = variant.Size
			variantData.Quantity = variant.Quantity
			_, err := s.variantRepositories.Update(variantData)
			if err != nil {
				continue
			}
		} else {
			create := entities.Variant{
				ProductID:  uuid.FromStringOrNil(variant.ProductID),
				Sku:        variant.Sku,
				Colour:     variant.Colour,
				Size:       variant.Size,
				Ingredient: variant.Ingredient,
				Quantity:   variant.Quantity,
			}
			_, errCreat := s.variantRepositories.Create(create)
			if errCreat != nil {
				break
			}
		}
	}
	return input.Variant, nil
}

func (s *varianService) checkDataExisting(productID string, input request.ProductRequestUpdate) bool {
	variantExistingDatas, _ := s.variantRepositories.FindVariantByProductID(productID)
	var isAvailable bool
	for _, exist := range variantExistingDatas {

		for _, dataInput := range input.Variant {
			if exist.ID.String() == dataInput.ID {
				isAvailable = true
				break
			}
		}
		if !isAvailable {
			err := s.variantRepositories.Delete(exist.ID.String())
			if err != nil {
				fmt.Println("delete")
			}
		}
	}
	return isAvailable
}

func (s *varianService) DeleteByIdProduct(productID string, loggedUser uuid.UUID) (bool, error) {
	variantExistingDatas, _ := s.variantRepositories.FindVariantByProductID(productID)
	for _, exist := range variantExistingDatas {
		exist.DeletedBy = loggedUser
		update, _ := s.variantRepositories.Update(exist)
		fmt.Println("update nya", update)
		delErr := s.variantRepositories.Delete(exist.ID.String())
		if delErr != nil {
			return false, delErr
		}
	}
	return true, nil
}
