package repositories

import (
	"ecommerce/entities"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu entities.Menu) (entities.Menu, error)
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
