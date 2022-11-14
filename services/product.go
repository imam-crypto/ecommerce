package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	"ecommerce/utils"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type ProductService interface {
	Create(loggedUser uuid.UUID, input request.ProductRequest) (entities.Product, error)
	FindProductByID(productID string) (entities.Product, error)
	Update(productID string, input request.ProductRequestUpdate) (entities.Product, error)
	Delete(id string, loggedUser uuid.UUID) (bool, error)
	GetAllProducts(searchFilter string, pagination utils.Pagination) ([]entities.Product, utils.Pagination)
}
type productService struct {
	productReposiories  repositories.ProductRepositories
	variantRepositories repositories.VariantRepository
}

func NewProductService(productReposiories repositories.ProductRepositories, variantRepositories repositories.VariantRepository) *productService {
	return &productService{productReposiories, variantRepositories}
}
func (s *productService) Create(loggedUser uuid.UUID, input request.ProductRequest) (entities.Product, error) {

	createProduct := entities.Product{
		Base: entities.Base{
			CreatedBy: loggedUser,
		},
		CategoryID:  uuid.Must(uuid.FromString(input.CategoryID)),
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
	}
	strore, errStore := s.productReposiories.Create(createProduct)
	if errStore != nil {
		return strore, errStore
	}
	return strore, nil
}

func (s *productService) FindProductByID(productID string) (entities.Product, error) {
	findProduct, errFind := s.productReposiories.FindProductByIDUpdate(productID)
	if errFind != nil {
		return findProduct, errFind
	}
	return findProduct, nil
}

func (s *productService) Update(productID string, input request.ProductRequestUpdate) (entities.Product, error) {
	oldProduct, errGet := s.productReposiories.FindProductByIDUpdate(productID)
	var isAvailable bool

	if errGet != nil {
		return oldProduct, errGet
	}
	for _, exist := range oldProduct.Variant {
		for _, dataInput := range input.Variant {
			if exist.ID.String() == dataInput.ID {
				isAvailable = true
				break
			}
		}
		if !isAvailable {
			//err := s.variantRepositories.Delete(exist.ID.String())
			//if err != nil {
			//	fmt.Println("delete")
			//}
			fmt.Println("[][][][] ini bedanya ", exist.ID.String())
		}
	}
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
				ProductID:  uuid.FromStringOrNil(productID),
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
	oldProduct.Title = input.Title
	oldProduct.Description = input.Description
	oldProduct.Price = input.Price
	//oldProduct.CategoryID = uuid.FromString(input.CategoryID)

	update, errUpdate := s.productReposiories.Update(oldProduct)
	fmt.Println("oold variantnyya", oldProduct.Variant)

	if errUpdate != nil {
		return update, errUpdate
	}
	return update, nil
}
func (s *productService) Delete(id string, loggedUser uuid.UUID) (bool, error) {
	getProduct, errGet := s.productReposiories.FindProductByIDUpdate(id)

	if errGet != nil {
		return false, errGet
	}
	getProduct.DeletedBy = loggedUser
	_, errUpdate := s.productReposiories.Update(getProduct)
	if errUpdate != nil {
		return false, errUpdate
	}
	errDel := s.productReposiories.Delete(id)
	if errDel != nil {
		return false, errDel
	}
	return true, nil
}

func (s *productService) GetAllProducts(searchFilter string, pagination utils.Pagination) ([]entities.Product, utils.Pagination) {
	query := ""
	if searchFilter != "" && query != "" {
		query += " AND LOWER(name) LIKE LOWER('%" + searchFilter + "%') AND LOWER(title) LIKE LOWER('%" + searchFilter + "%')"
	} else if searchFilter != "" && query == "" {
		query += "LOWER(title) LIKE LOWER('%" + searchFilter + "%') OR LOWER(title) LIKE LOWER('%" + searchFilter + "%')"
	}
	products, pagination := s.productReposiories.GetAllProducts(pagination, query)
	return products, pagination
}
