package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string                   `json:"message"`
	Details []ValidationErrorDetails `json:"details,omitempty"`
}

type ValidationErrorDetails struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ResponseOK(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusOK, Response{Message: appconstant.MsgOK, Data: res})
}

func ResponseCreated(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, Response{Message: appconstant.MsgCreated, Data: res})
}

func ResponseRegister(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, Response{Message: appconstant.MsgCheckEmailRegister, Data: res})
}

func ResponseVerifRegister(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, Response{Message: appconstant.MsgCheckEmailRegister, Data: res})
}

func ResponseLogin(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, Response{Message: appconstant.MsgCheckEmailLogin, Data: res})
}