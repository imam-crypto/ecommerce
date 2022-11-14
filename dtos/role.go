package dtos

import "ecommerce/entities"

type ResponseRole struct {
	ID         string               `json:"id"`
	Name       string               `json:"name"`
	MenuAccess []MenuAccessResponse `json:"menu_access"`
}

func ConvResponseRole(role entities.Role) ResponseRole {
	conv := ResponseRole{
		ID:   role.ID.String(),
		Name: role.Name,
	}
	for _, menuObj := range role.MenuAccess {
		menuRes := ConvMenuAccessResponse(menuObj)
		conv.MenuAccess = append(conv.MenuAccess, menuRes)
	}
	return conv
}
