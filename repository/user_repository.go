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
	CreateOTP(ctx context.Context, userId string) (*string, error)
	VerifyOTP(ctx context.Context, userId int, otp string) (bool, error)
	UpdateUserVerificationStatus(ctx context.Context, userId int) error
	CreateRoleUser(ctx context.Context, userId int)  error
}

type userRepositoryDB struct {
	db DBTX
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{
		db: db,
	}
}

func (r *userRepositoryDB) FindAccountByOtp(ctx context.Context, otp string) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByEmailQuery, otp).Scan(
		&account.Id, &account.Photo, &account.FirstName, &account.LastName, &account.Username, &account.Email, &account.Gender, &account.Address, &account.PhoneNumber, &account.Password, &account.EmailVerifiedAt, &account.RoleId, &account.RoleName, &account.RoleCode)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &account, nil
}

func (r *userRepositoryDB) FindAccountByEmail(ctx context.Context, email string) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByEmailQuery, email).Scan(
		&account.Id, 
		&account.Photo, 
		&account.FirstName, 
		&account.LastName, 
		&account.Username, 
		&account.Email, 
		&account.Gender, 
		&account.Address, 
		&account.PhoneNumber, 
		&account.Password, 
		&account.EmailVerifiedAt, 
		&account.RoleId, 
		&account.RoleName, 
		&account.RoleCode)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &account, nil
}

func (r *userRepositoryDB) FindAccountByUsername(ctx context.Context, username string) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByUsernameQuery, username).Scan(
		&account.Id, 
		&account.Photo, 
		&account.FirstName, 
		&account.LastName, 
		&account.Username, 
		&account.Email, 
		&account.Gender, 
		&account.Address, 
		&account.PhoneNumber, 
		&account.Password, 
		&account.EmailVerifiedAt, 
		&account.RoleId, 
		&account.RoleName, 
		&account.RoleCode)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &account, nil
}

func (r *userRepositoryDB) CreateOTP(ctx context.Context, userId string) (*string, error) {

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := fmt.Sprintf("%06d", rng.Intn(900000)+100000)

	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	otpExpireTime := time.Now().In(loc).Add(10 * time.Minute)

	otpExpireTimeUTC := otpExpireTime.UTC().Add(7 * time.Hour)

	_, err = r.db.ExecContext(ctx, database.InserOtpQuery, userId, otp, otpExpireTimeUTC)
	if err != nil {
		return nil, err
	}
	return &otp, nil
}

func (r *userRepositoryDB) PostOneUser(ctx context.Context, user entity.User) (*int, error) {
	result, err := r.db.ExecContext(ctx, database.PostOneAccountQuery,
		user.Email,
		user.Gender,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	userIdInt := int(userId)
	return &userIdInt, nil
}

func (r *userRepositoryDB) CreateRoleUser(ctx context.Context, userId int)  error {
	roles_id:= 2
	_, err := r.db.ExecContext(ctx, database.PostRoleUserQuery,
		userId,
		roles_id,
	)
	if err != nil {
		return err
	}

	return nil
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

func (r *userRepositoryDB) VerifyOTP(ctx context.Context, userId int, otp string) (bool, error) {
	var otpRecord string
	err := r.db.QueryRowContext(ctx, database.FindUserOtp, userId, otp).Scan(&otpRecord)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil 
		}
		return false, err // Some other error occurred
	}

	return true, nil // OTP is valid
}

func (r *userRepositoryDB) UpdateUserVerificationStatus(ctx context.Context, userId int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET email_verified_at = NOW() WHERE id = ?", userId)
	return err
}
