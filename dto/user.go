package dto

import "github.com/sidiqPratomo/DJKI-Pengaduan/entity"

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email" validate:"required,email"`
	Username  string `json:"username" binding:"required" validate:"required"`
	First_name string `json:"first_name" binding:"required" validate:"required"`
    Last_name string `json:"last_name" binding:"required"`
    Password string `json:"password" binding:"required"`
    Password_confirmation string `json:"password_confirmation" binding:"required"`
    Gender int `json:"gender" binding:"required"`
    Phone_number string `json:"phone_number" binding:"required"`
}

func RegisterRequestToAccount(RegisterRquest RegisterRequest) entity.User {
	return entity.User{
		Email: RegisterRquest.Email,
		Username:  RegisterRquest.Username,
	}
}