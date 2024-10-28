package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
)

func ResponseOK(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusOK, dto.Response{Message: appconstant.MsgOK, Data: res})
}

func ResponseCreated(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, dto.Response{Message: appconstant.MsgCreated, Data: res})
}

func ResponseRegister(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusCreated, dto.Response{Message: appconstant.MsgCheckEmailRegister, Data: res})
}