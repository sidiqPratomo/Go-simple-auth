package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
	"github.com/sidiqPratomo/DJKI-Pengaduan/usecase"
)

type AuthenticationHandler struct {
	authenticationUsecase usecase.AuthenticationUsecase
}

func NewAuthenticationHandler(authenticationUsecase usecase.AuthenticationUsecase) AuthenticationHandler {
	return AuthenticationHandler{
		authenticationUsecase: authenticationUsecase,
	}
}

func (h *AuthenticationHandler) RegisterUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.authenticationUsecase.RegisterUser(ctx.Request.Context(), registerRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseRegister(ctx, nil)
}
