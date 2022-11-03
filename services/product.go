package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

type ProductService interface {
	Create(loggedUser uuid.UUID, input request.ProductRequest) (entities.Product, error)
	FindProductByID(productID string) (entities.Product, error)
}
type productService struct {
	productReposiories repositories.ProductRepositories
}

func NewProductService(productReposiories repositories.ProductRepositories) *productService {
	return &productService{productReposiories}
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
	findProduct, errFind := s.productReposiories.FindProductByID(productID)
	if errFind != nil {
		return findProduct, errFind
	}
	return findProduct, nil
}
