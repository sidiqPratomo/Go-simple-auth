package entity

import "time"

type User struct {
	Id              int64
	Photo           string
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Gender          string
	Address         string
	PhoneNumber     string
	Password        string
	EmailVerifiedAt *time.Time
	RememberToken   string
	CreatedBy       string
	UpdatedBy       string
	CreatedTime     *time.Time
	UpdatedTime     *time.Time
	Status          int8
}

type Roles struct {
	Id          int64
	Name        string
	Code        string
	CreatedBy   string
	UpdatedBy   string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Status      int8
}

type RoleUsers struct {
	Id          int64
	UserId      int
	RolesId     int
	CreatedBy   string
	UpdatedBy   string
	CreatedTime *time.Time
	UpdatedTime *time.Time
	Status      int8
}

type Priveleges struct{
	Id int64
	Module string
	Submodule string
	Ordering string
	Action string
	Method string
	Uri	string
	CreatedBy       string
	UpdatedBy       string
	CreatedTime     *time.Time
	UpdatedTime     *time.Time
	Status          int8
}

type RolePrivileges struct{
	Id int
	Role int16
	Action string
	Method string
	Uri	string
	CreatedBy       string
	UpdatedBy       string
	CreatedTime     *time.Time
	UpdatedTime     *time.Time
	Status          int8
}
