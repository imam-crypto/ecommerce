package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type ProductRepositories interface {
	Create(product entities.Product) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
	FindProductByID(productID string) (entities.Product, error)
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
func (r *productRepositories) Update(product entities.Product) (entities.Product, error) {

	errUpdate := r.db.Save(&product).Error

	if errUpdate != nil {
		return product, errUpdate
	}
	return product, errUpdate
}
