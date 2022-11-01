package formatresponse

import "ecommerce/entities"

type ResponseUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func PaginateUserResponse(us *entities.User) ResponseUser {
	FormatUser := ResponseUser{
		ID:       us.Base.ID.String(),
		Username: us.Username,
		Email:    us.Email,
		Role:     us.Role,
	}
	return FormatUser

}

func ConvResponseUser(user entities.User) ResponseUser {
	formatterUser := ResponseUser{}
	formatterUser.Username = user.Username
	formatterUser.Email = user.Email
	formatterUser.Role = user.Role
	return formatterUser
}
