package dtos

import (
	"ecommerce/entities"
	uuid "github.com/satori/go.uuid"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Role     ResponseRole `json:"role"`
}

type UpdateRequest struct {
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Gender     string `json:"gender"`
	Name       string `json:"name"`
}
type LoginResponse struct {
	ID           uuid.UUID    `json:"id"`
	Email        string       `json:"email"`
	Role         ResponseRole `json:"role"`
	Token        string       `json:"token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    string       `json:"expire_at"`
}

func ConvRegisterResponse(us entities.User) RegisterResponse {
	FormatUser := RegisterResponse{
		ID:       us.Base.ID.String(),
		Username: us.Username,
		Email:    us.Email,
		Role:     ConvResponseRole(us.Role),
	}
	return FormatUser
}

func ConvLoginResponse(user entities.User, token, refreshToken, ExpiresAt string) LoginResponse {
	formatter := LoginResponse{
		ID:           user.ID,
		Email:        user.Email,
		Role:         ConvResponseRole(user.Role),
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresAt:    ExpiresAt,
	}
	return formatter
}
