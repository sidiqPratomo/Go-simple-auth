package repository

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/sidiqPratomo/DJKI-Pengaduan/database"
	"github.com/sidiqPratomo/DJKI-Pengaduan/entity"
)

type UserRepository interface {
	PostOneUser(ctx context.Context, account entity.User) (*int, error)
	GetAllUser(ctx context.Context, user entity.User) ([]entity.User, error)
	FindAccountByEmail(ctx context.Context, email string) (*entity.UserRoles, error)
	FindAccountByUsername(ctx context.Context, username string) (*entity.UserRoles, error)
	CreateOTP(ctx context.Context, userId string)(*string, error)
}

type userRepositoryDB struct {
	db DBTX
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{
		db: db,
	}
}

func (r *userRepositoryDB) FindAccountByEmail(ctx context.Context, email string) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByEmailQuery, email).Scan(
		&account.Id, &account.Photo, &account.FirstName, &account.LastName, &account.Username, &account.Email, &account.Gender, &account.Address, &account.PhoneNumber, &account.Password, &account.EmailVerifiedAt, &account.RoleId, &account.RoleName, &account.RoleCode)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *userRepositoryDB) FindAccountByUsername(ctx context.Context, username string) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByEmailQuery, username).Scan(
		&account.Id, &account.Photo, &account.FirstName, &account.LastName, &account.Username, &account.Email, &account.Gender, &account.Address, &account.PhoneNumber, &account.Password, &account.EmailVerifiedAt, &account.RoleId, &account.RoleName, &account.RoleCode)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *userRepositoryDB) CreateOTP(ctx context.Context, userId string)(*string, error) {
	
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := fmt.Sprintf("%06d", rng.Intn(900000)+100000)

	otpExpireTime := time.Now().Add(10 * time.Minute)

	_, err:= r.db.ExecContext(ctx, database.InserOtpQuery,userId, otp, otpExpireTime)
	if err!=nil{
		return nil, err
	}
	return &otp, nil
}

func (r *userRepositoryDB) PostOneUser(ctx context.Context, user entity.User) (*int, error) {
	var userId int

	err := r.db.QueryRowContext(ctx, database.PostOneAccountQuery,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
	).Scan(&userId)
	if err != nil {
		return nil, err
	}

	return &userId, nil
}

func (r *userRepositoryDB) GetAllUser(ctx context.Context, user entity.User) ([]entity.User, error) {
	rows, err := r.db.QueryContext(ctx, database.GetAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.Id,
			&user.Photo,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Gender,
			&user.Address,
			&user.PhoneNumber,
			&user.EmailVerifiedAt,
			&user.RememberToken,
			&user.CreatedBy,
			&user.UpdatedBy,
			&user.CreatedTime,
			&user.UpdatedTime,
			&user.Status,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
