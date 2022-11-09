package dtos

import "ecommerce/entities"

type ResponseRole struct {
	Name string
}

func ConvResponseRole(role entities.Role) ResponseRole {
	conv := ResponseRole{
		Name: role.Name,
	}
	return conv
}
