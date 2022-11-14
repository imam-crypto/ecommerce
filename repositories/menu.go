package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu entities.Menu) (entities.Menu, error)
	Delete(id string) error
}
type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) Create(menu entities.Menu) (entities.Menu, error) {
	errCreate := r.db.Create(&menu).Error
	if errCreate != nil {
		return menu, errCreate
	}
	return menu, nil
}
func (r *menuRepository) Delete(id string) error {
	deleteErr := r.db.Where("id=?", id).Delete(&entities.Menu{}).Error
	return deleteErr
}
