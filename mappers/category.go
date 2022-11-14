package mappers

import (
	"ecommerce/entities"
	"ecommerce/request"
	uuid "github.com/satori/go.uuid"
)

func CreateCategory(loggedUser uuid.UUID, input request.CategoryRequestInsert, url string) entities.Category {
	mapCreate := entities.Category{
		Base: entities.Base{
			CreatedBy: loggedUser,
		},
		Name:          input.Name,
		PublicIDCloud: input.PublicIDCloud,
		UrlImage:      url,
	}
	return mapCreate
}

func UpdateCategory(oldValue entities.Category, newValue request.CategoryRequestInsert, url string) entities.Category {
	var newCode = newValue.PublicIDCloud
	if newCode != "" {
		oldValue.PublicIDCloud = newCode
	}
	oldValue.Name = newValue.Name
	oldValue.UrlImage = url

	return oldValue
}
