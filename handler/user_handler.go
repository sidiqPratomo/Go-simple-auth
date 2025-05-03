package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
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

func (h *UserHandler) ReadUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	user, err := h.userUsecase.ReadUser(ctx, idInt)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBy, exists := ctx.Get(appconstant.AccountId)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user.Id = idInt
	updatedByStr := fmt.Sprintf("%v", updatedBy)
	user.UpdatedBy = &updatedByStr
	result, err := h.userUsecase.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": result})
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	updatedBy, exists := ctx.Get(appconstant.AccountId)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	updatedByStr := fmt.Sprintf("%v", updatedBy)

	err = h.userUsecase.SoftDeleteUser(ctx, idInt, updatedByStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}