package dto

import "github.com/sidiqPratomo/DJKI-Pengaduan/entity"

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email" validate:"required,email"`
	Username  string `json:"name" binding:"required" validate:"required"`
}

func RegisterRequestToAccount(RegisterRquest RegisterRequest) entity.User {
	return entity.User{
		Email: RegisterRquest.Email,
		Username:  RegisterRquest.Username,
	}
}