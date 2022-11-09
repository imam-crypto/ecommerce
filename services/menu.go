package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

type MenuService interface {
	Create(loggedUser uuid.UUID, input request.MenuRequest) (bool, error)
}

type menuService struct {
	serviceRepo repositories.MenuRepository
}

func NewMenuService(serviceRepo repositories.MenuRepository) *menuService {
	return &menuService{serviceRepo}
}

func (s *menuService) Create(loggedUser uuid.UUID, input request.MenuRequest) (bool, error) {
	mapCreate := entities.Menu{
		Base: entities.Base{
			CreatedBy: loggedUser,
		},
		MenuName: input.MenuName,
		Uri:      input.Uri,
	}
	_, err := s.serviceRepo.Create(mapCreate)
	if err != nil {
		return false, err
	}
	return true, nil
}
