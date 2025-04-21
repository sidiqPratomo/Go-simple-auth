package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/dto"
	"github.com/sidiqPratomo/DJKI-Pengaduan/usecase"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) IndexUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	// Ambil raw string dari query dan konversi
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	skip, _ := strconv.Atoi(ctx.DefaultQuery("skip", ""))
	statusStr := ctx.DefaultQuery("status", "")
	var statusPtr *int8 = nil
	if statusStr != "" {
		s, err := strconv.Atoi(statusStr)
		if err == nil {
			tmp := int8(s)
			statusPtr = &tmp
		}
	}

	rawParams := util.QueryParam{
		Offset:    int32(skip),
		SortBy:    ctx.DefaultQuery("sortBy", ""),
		SortOrder: ctx.DefaultQuery("sort", ""),
		Page:      page,
		Limit:     int32(limit),
		Status:    statusPtr,
	}
	fmt.Println("rawParams", rawParams.Offset)
	params, err := util.SetDefaultQueryParams(rawParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	dtoQueryParams := dto.UserQueryParams{
		Limit:     params.Limit,
		Offset:    params.Offset,
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
		Status:    statusPtr,
	}

	users, err := h.userUsecase.IndexUser(ctx, dtoQueryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
