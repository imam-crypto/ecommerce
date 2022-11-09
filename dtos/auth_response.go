package dtos

import (
	"ecommerce/entities"
	uuid "github.com/satori/go.uuid"
)

type ResponseUserLogin struct {
	ID           uuid.UUID    `json:"id"`
	Email        string       `json:"email"`
	Role         ResponseRole `json:"role"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    string       `json:"expire_at"`
}

func ConvResponseUserLogin(user entities.User, token, refreshToken, ExpiresAt string) ResponseUserLogin {
	formatter := ResponseUserLogin{
		ID:           user.ID,
		Email:        user.Email,
		Role:         ConvResponseRole(user.Role),
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    ExpiresAt,
	}
	return formatter
}
