package repositories

import (
	"ecommerce/entities"
	"ecommerce/utils"
	"gorm.io/gorm"
)

type ProductRepositories interface {
	GetAllProducts(pagination utils.Pagination, queryFilter string) ([]entities.Product, utils.Pagination)
	Create(product entities.Product) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
	FindProductByID(productID string) (entities.Product, error)
	FindProductByIDUpdate(productID string) (entities.Product, error)
	Delete(id string) error
}
type productRepositories struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepositories {
	return &productRepositories{db}
}
func (r *productRepositories) Create(product entities.Product) (entities.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *productRepositories) FindProductByID(productID string) (entities.Product, error) {
	var product entities.Product
	errFind := r.db.Preload("Variant").Where("id =?", productID).Find(&product).Error
	if errFind != nil {
		return product, errFind
	}
	return product, nil
}
func (r *productRepositories) FindProductByIDUpdate(productID string) (entities.Product, error) {
	var product entities.Product
	errFind := r.db.Where("id =?", productID).Find(&product).Error
	if errFind != nil {
		return product, errFind
	}
	return product, nil
}
func (r *productRepositories) Update(product entities.Product) (entities.Product, error) {

	errUpdate := r.db.Save(&product).Error

	if errUpdate != nil {
		return product, errUpdate
	}
	return product, errUpdate
}

func (r *productRepositories) Delete(id string) error {
	var (
		product entities.Product
		//variant entities.Variant
	)
	deleteErr := r.db.Where("id =?", id).Delete(&product).Error
	//delVariant := r.db.Where("product_id =? ", id).Delete(&variant).Error
	return deleteErr
}
func (r *productRepositories) GetAllProducts(pagination utils.Pagination, queryFilter string) ([]entities.Product, utils.Pagination) {
	var products []entities.Product

	r.db.Scopes(utils.Paginate(&products, &pagination, r.db)).Preload("Category").Preload("Variant").Where(queryFilter).Find(&products)

	return products, pagination
}
