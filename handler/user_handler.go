package handler

import (
	"net/http"
	"strconv"
	"strings"

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
	limit, _ := strconv.Atoi(ctx.DefaultQuery("!limit", "10"))
	skip, _ := strconv.Atoi(ctx.DefaultQuery("!skip", ""))
	statusStr := ctx.DefaultQuery("status", "")

	var statusPtr *int8 = nil
	if statusStr != "" {
		s, err := strconv.Atoi(statusStr)
		if err == nil {
			tmp := int8(s)
			statusPtr = &tmp
		}
	}

	sortBy, sortOrder := parseSortQuery(ctx)

	rawParams := util.QueryParam{
		Offset:    int32(skip),
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Limit:     int32(limit),
		Status:    statusPtr,
	}

	dtoQueryParams := dto.UserQueryParams{
		Limit:     rawParams.Limit,
		Offset:    rawParams.Offset,
		SortBy:    rawParams.SortBy,
		SortOrder: rawParams.SortOrder,
		Status:    statusPtr,
	}

	users, err := h.userUsecase.IndexUser(ctx, dtoQueryParams)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func parseSortQuery(ctx *gin.Context) (string, string) {
	query := ctx.Request.URL.Query()

	for key, values := range query {
		if strings.HasPrefix(key, "!sort[") && strings.HasSuffix(key, "]") && len(values) > 0 {
			field := key[6 : len(key)-1]
			direction := values[0]
			if direction == "-1" {
				return field, "DESC"
			} else if direction == "1" {
				return field, "ASC"
			}
		}
	}
	return "id", "DESC" // default fallback
}