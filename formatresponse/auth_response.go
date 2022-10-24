package formatresponse

import (
	"ecommerce/entities"
	uuid "github.com/satori/go.uuid"
)

type ResponseUserLogin struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Token     string    `json:"token"`
	ExpiresAt string    `json:"expire_at"`
}

func ConvResponseUserLogin(user entities.User, token, ExpiresAt string) ResponseUserLogin {
	formatter := ResponseUserLogin{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		Token:     token,
		ExpiresAt: ExpiresAt,
	}
	return formatter
}
