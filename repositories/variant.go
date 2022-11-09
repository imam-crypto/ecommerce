package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type VariantRepository interface {
	Create(variant entities.Variant) (entities.Variant, error)
	FindVariantByProductID(productID string) ([]entities.Variant, error)
	Update(variant entities.Variant) (entities.Variant, error)
	FindVariantByID(variantID string) (entities.Variant, error)
	Delete(id string) error
}
type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) *variantRepository {
	return &variantRepository{db}
}
func (r *variantRepository) Create(variant entities.Variant) (entities.Variant, error) {
	errCreate := r.db.Create(&variant).Error
	if errCreate != nil {
		return variant, errCreate
	}
	return variant, nil
}
func (r *variantRepository) FindVariantByProductID(productID string) ([]entities.Variant, error) {
	var variants []entities.Variant

	err := r.db.Where("product_id =?", productID).Find(&variants).Error
	if err != nil {
		return variants, err
	}

	return variants, nil
}
func (r *variantRepository) FindVariantByID(variantID string) (entities.Variant, error) {
	var variants entities.Variant

	err := r.db.Where("id =?", variantID).Find(&variants).Error
	if err != nil {
		return variants, err
	}

	return variants, nil
}
func (r *variantRepository) Update(variant entities.Variant) (entities.Variant, error) {

	errUpdate := r.db.Save(&variant).Error

	if errUpdate != nil {
		return variant, errUpdate
	}
	return variant, errUpdate
}
func (r *variantRepository) Delete(id string) error {
	deleteErr := r.db.Where("id=?", id).Delete(&entities.Variant{}).Error
	return deleteErr
}
