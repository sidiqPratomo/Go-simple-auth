package dto

import (
	"errors"
	"time"

	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"
	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/DJKI-Pengaduan/entity"
	"github.com/sidiqPratomo/DJKI-Pengaduan/util"
)

type RegisterRequest struct {
	Nik                   string `json:"nik" binding:"required" validate:"required"`
	Email                 string `json:"email" binding:"required,email" validate:"required,email"`
	Username              string `json:"username" binding:"required" validate:"required"`
	First_name            string `json:"first_name" binding:"required" validate:"required"`
	Last_name             string `json:"last_name" binding:"required"`
	Password              string `json:"password" binding:"required"`
	Password_confirmation string `json:"password_confirmation" binding:"required"`
	Gender                int    `json:"gender" binding:"required"`
	Phone_number          string `json:"phone_number" binding:"required"`
}

func RegisterRequestToAccount(RegisterRquest RegisterRequest) (entity.User, error) {
	var Gender string
	if RegisterRquest.Gender == 1 {
		Gender = "Male"
	} else {
		Gender = "Female"
	}
	if RegisterRquest.Password != RegisterRquest.Password_confirmation {
		return entity.User{}, apperror.BadRequestError(errors.New("passwords do not match"))
	}
	isNameValid := util.RegexValidate(RegisterRquest.Username, appconstant.NameRegexPattern)
	if !isNameValid {
		return entity.User{}, apperror.InvalidNameError(errors.New("invalid name"))
	}
	return entity.User{
		Nik:         &RegisterRquest.Nik,
		Email:       RegisterRquest.Email,
		Username:    RegisterRquest.Username,
		FirstName:   RegisterRquest.First_name,
		LastName:    RegisterRquest.Last_name,
		Password:    RegisterRquest.Password,
		Gender:      &Gender,
		PhoneNumber: &RegisterRquest.Phone_number,
	}, nil
}

type VerifyOTPRequest struct {
	Username string `json:"username" binding:"required"`
	OTP      string `json:"otp" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VerifyUserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	OTP string `json:"otp" binding:"required"`
}

type User struct {
	Id              int64      `json:"id"`
	StatusOTP		*int8      `json:"status_otp"`
	Nik				*string	   `json:"nik"`
	Photo           *string    `json:"photo"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Gender          *string     `json:"gender"`
	Address         *string     `json:"address"`
	PhoneNumber     *string     `json:"phone_number"`
	EmailVerifiedAt *time.Time  `json:"email_verified_at"`
	CreatedBy       *string    `json:"created_by"`
	UpdatedBy       *string    `json:"updated_by"`
	CreatedTime     *time.Time  `json:"created_time"`
	UpdatedTime     *time.Time  `json:"updated_time"`
	Status          int        `json:"status"`
	Role            []string   `json:"role"`
	Roles           []UserRole `json:"roles"`
}

type UserDetail struct {
	Id              int64      `json:"id"`
	Nik				*string	   `json:"nik"`
	Photo           *string    `json:"photo"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Gender          *string     `json:"gender"`
	Address         *string     `json:"address"`
	PhoneNumber     *string     `json:"phone_number"`
	EmailVerifiedAt *time.Time  `json:"email_verified_at"`
	CreatedBy       *string    `json:"created_by"`
	UpdatedBy       *string    `json:"updated_by"`
	CreatedTime     *time.Time  `json:"created_time"`
	UpdatedTime     *time.Time  `json:"updated_time"`
	Status          int        `json:"status"`
}

type UserQueryParams struct{
	Limit int32
	Offset int32
	SortBy string
	SortOrder string
	Status *int8
}

type Users struct{
	Users []User `json:"users"`
}

type ResponseIndex[T any] struct {
	Status  bool   `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type PagedResult[T any] struct {
	Result []T `json:"result"`
	Count  int `json:"count"`
}

func MapDTOQuerytoEntity(params UserQueryParams) entity.UserQuery {
	var status *int
	if params.Status != nil {
		converted := int(*params.Status)
		status = &converted
	}
	return entity.UserQuery{
		Limit:     int(params.Limit),
		Offset:    int(params.Offset),
		SortBy:    params.SortBy,
		SortOrder: params.SortOrder,
		Status:    status,
	}
}

// UserRole represents the structure of roles associated with the user.
type UserRole struct {
	Id          int64      `json:"id"`
	UsersId     int64      `json:"users_id"`
	RolesId     RoleDetail `json:"roles_id"`
	CreatedBy   *string    `json:"created_by"`
	UpdatedBy   *string    `json:"updated_by"`
	CreatedTime *string  `json:"created_time"`
	UpdatedTime *string  `json:"updated_time"`
	Status      int        `json:"status"`
}

// RoleDetail represents detailed information about each role.
type RoleDetail struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	CreatedBy   *string   `json:"created_by"`
	UpdatedBy   *string   `json:"updated_by"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	Status      int       `json:"status"`
}

// Privilege represents each privilege associated with a role.
type Privilege struct {
	Id          int64     `json:"id"`
	Role        int64     `json:"role"`
	Action      string    `json:"action"`
	Uri         string    `json:"uri"`
	Method      string    `json:"method"`
	CreatedBy   *string   `json:"created_by"`
	UpdatedBy   *string   `json:"updated_by"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
	Status      int       `json:"status"`
}

// Role represents the structure of the role in the response.
type Role struct {
	Privileges []Privilege `json:"privileges"`
	Role       []UserRole  `json:"role"`
}

// VerifyLoginUserResponse represents the full response structure for VerifyLoginUser.
type VerifyLoginUserResponse struct {
	User        User   `json:"user"`
	Role        Role   `json:"role"`
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type LoginResponse struct {
	User        User   `json:"user"`
	Role        Role   `json:"role"`
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

func MapRolesToDTOs(roles []entity.RoleUsers) []UserRole {
	var roleDTOs []UserRole
	for _, role := range roles {
		roleDTOs = append(roleDTOs, UserRole{
			Id:      role.Id,
			UsersId:  int64(role.UserId),
			RolesId: RoleDetail{
				Id:        role.RolesId.Id,
				Name:      role.RolesId.Name,
				Code:      role.RolesId.Code,
				CreatedBy: role.RolesId.CreatedBy,
				UpdatedBy: role.RolesId.UpdatedBy,
				Status:    int(role.RolesId.Status),
			},
			CreatedBy:  role.CreatedBy,
			UpdatedBy:  role.UpdatedBy,
			CreatedTime: role.CreatedTime,
			UpdatedTime: role.UpdatedTime,
			Status:     int(role.Status),
		})
	}
	return roleDTOs
}

func MapPrivilegesToDTOs(privileges []entity.RolePrivileges) []Privilege {
	var privilegeDTOs []Privilege
	for _, priv := range privileges {
		privilegeDTOs = append(privilegeDTOs, Privilege{
			Id:        int64(priv.Id),
			Role:      int64(priv.Role),
			Action:    priv.Action,
			Uri:       priv.Uri,
			Method:    priv.Method,
			CreatedBy: priv.CreatedBy,
			UpdatedBy: priv.UpdatedBy,
			Status:    int(priv.Status),
		})
	}
	return privilegeDTOs
}
