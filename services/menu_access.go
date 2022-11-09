package services

import (
	"ecommerce/entities"
	"ecommerce/repositories"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

type MenuAccessService interface {
	CreateMenuAccess(loggedUser uuid.UUID, input request.MenuAccessRequest) (bool, error)
}

type menuAccessService struct {
	menuAccessRepository repositories.MenuAccessRepository
}

func NewMenuAccessService(menuAccessRepository repositories.MenuAccessRepository) *menuAccessService {
	return &menuAccessService{menuAccessRepository}
}

func (s *menuAccessService) CreateMenuAccess(loggedUser uuid.UUID, input request.MenuAccessRequest) (bool, error) {
	//var menus = []entities.MenuAccess{}
	var result bool
	for _, menuID := range input.MenuID {
		mapCreateMenuAccess := entities.MenuAccess{
			Base: entities.Base{
				CreatedBy: loggedUser,
			},
			RoleID:       uuid.FromStringOrNil(input.RoleID),
			MenuID:       uuid.FromStringOrNil(menuID),
			ReadAccess:   input.ReadAccess,
			CreateAccess: input.CreateAccess,
			UpdateAccess: input.UpdateAccess,
			DeleteAccess: input.DeleteAccess,
		}
		//menus = append(menus,menuID)menuID
		createAccess, errCreate := s.menuAccessRepository.Create(mapCreateMenuAccess)
		if errCreate != nil {
			return result, errCreate
		}
		return createAccess, nil
	}
	return result, nil
}
