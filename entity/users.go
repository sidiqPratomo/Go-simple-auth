package entity

import (
	"time"
)

type UserOtps struct {
	Id          int64
	User_id     int64
	Otp         string
	Expired_at  *time.Time
	CreatedBy   string
	UpdatedBy   string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Status      int8
}

type UserQuery struct {
	Limit  int
	Offset int
	SortBy string
	SortOrder string
	Status *int
}

type UserRoles struct {
	Id              int64
	StatusOTP		int8
	Nik				*string
	Photo           *string
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Gender          *string
	Address         *string
	PhoneNumber     *string
	Password        string
	EmailVerifiedAt *time.Time
	RoleId          int64
	RoleName        string
	RoleCode        string
	Status			int
}

type User struct {
	Id              int64
	StatusOTP		int8
	Nik				*string
	Photo           *string
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Gender          *string
	Address         *string
	PhoneNumber     *string
	Password        string
	EmailVerifiedAt *time.Time
	RememberToken   *string
	CreatedBy       *string
	UpdatedBy       *string
	CreatedTime     *time.Time
	UpdatedTime     *time.Time
	Status          int8
}

type Roles struct {
	Id          int64
	Name        string
	Code        string
	CreatedBy   *string
	UpdatedBy   *string
	CreatedTime *string
	UpdatedTime *string
	Status      int8
}

type RoleUsers struct {
	Id          int64
	UserId      int
	RolesId     Roles
	CreatedBy   *string
	UpdatedBy   *string
	CreatedTime *string
	UpdatedTime *string
	Status      int8
}

type Priveleges struct {
	Id          int64
	Module      string
	Submodule   string
	Ordering    string
	Action      string
	Method      string
	Uri         string
	CreatedBy   *string
	UpdatedBy   *string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Status      int8
}

type RolePrivileges struct {
	Id          int
	Role        int16
	Action      string
	Method      string
	Uri         string
	CreatedBy   *string
	UpdatedBy   *string
	CreatedTime *string
	UpdatedTime *string
	Status      int8
}
