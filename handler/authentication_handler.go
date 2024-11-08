package handler

import (
	"net/http"

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

func (h *AuthenticationHandler) VerifyLoginUser(ctx *gin.Context){
	ctx.Header("Content-Type", "application/json")

	var verifyLoginRequest dto.VerifyUserLoginRequest

	err := ctx.ShouldBindJSON(&verifyLoginRequest)
	if err != nil{
		ctx.Error(err)
		return
	}

	response, err := h.authenticationUsecase.VerifyUserLogin(ctx, verifyLoginRequest)
	if err != nil{
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *AuthenticationHandler) Login(ctx *gin.Context){
	ctx.Header("Content-Type", "application/json")

	var loginRequest dto.LoginRequest

	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil{
		ctx.Error(err)
		return
	}

	err = h.authenticationUsecase.LoginUser(ctx, loginRequest)
	if err != nil{
		ctx.Error(err)
		return
	}

	dto.ResponseLogin(ctx, nil)
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

func (h *AuthenticationHandler) VerifyRegisterUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var verifRegisterRequest dto.VerifyOTPRequest

	err := ctx.ShouldBindJSON(&verifRegisterRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = h.authenticationUsecase.VerifyUserRegister(ctx.Request.Context(), verifRegisterRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	dto.ResponseVerifRegister(ctx, nil)
}
