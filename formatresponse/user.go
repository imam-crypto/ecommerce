package formatresponse

import "ecommerce/entities"

type ResponseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func ConvResponseUser(user entities.User) ResponseUser {
	formatterUser := ResponseUser{}
	formatterUser.Username = user.Username
	formatterUser.Email = user.Email
	formatterUser.Role = user.Role
	return formatterUser
}
