package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category entities.Category) (entities.Category, error)
	FindByID(id string) (entities.Category, error)
	Update(category entities.Category) (entities.Category, error)
	Delete(id string) error
}
type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}
func (r *categoryRepository) Create(category entities.Category) (entities.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}
func (r *categoryRepository) FindByID(id string) (entities.Category, error) {
	var category entities.Category
	err := r.db.Where("id = ?", id).Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) Update(category entities.Category) (entities.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id string) error {
	deleteErr := r.db.Where("id=?", id).Delete(&entities.Category{}).Error
	return deleteErr
}
