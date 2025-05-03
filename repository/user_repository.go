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
	UpdateUser(ctx context.Context, user entity.User) error
	FindByID(ctx context.Context, id int) (*entity.User, error)
	FindAll(ctx context.Context, params entity.UserQuery) ([]entity.User, int, error)
	PostOneUser(ctx context.Context, account entity.User) (*int, error)
	GetAllUser(ctx context.Context, user entity.User) ([]entity.User, error)
	FindAccountByEmail(ctx context.Context, email string) (*entity.UserRoles, error)
	FindAccountByUsername(ctx context.Context, username string) (*entity.UserRoles, error)
	CreateOTP(ctx context.Context, userId string) (*string, error)
	VerifyOTP(ctx context.Context, userId int, otp string) (bool, error)
	UpdateUserVerificationStatus(ctx context.Context, userId int) error
	CreateRoleUser(ctx context.Context, userId int) error
	GetOTPByCode(ctx context.Context, otp string, userId int) (*entity.UserOtps, error)
	FindAccountByUserId(ctx context.Context, userId int) (*entity.UserRoles, error)
	GetUserRoles(ctx context.Context, userId int64) ([]entity.RoleUsers, []entity.RolePrivileges, error)
}

type userRepositoryDB struct {
	db DBTX
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{
		db: db,
	}
}

func (r *userRepositoryDB) UpdateUser(ctx context.Context, user entity.User) error {
	query := database.UpdateUserQuery

	_, err := r.db.ExecContext(ctx, query,
		user.StatusOTP,
		user.Nik,
		user.Photo,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.Gender,
		user.Address,
		user.PhoneNumber,
		user.Status,
		user.UpdatedBy,
		user.Id,
	)
	return err
}

func (r *userRepositoryDB) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var account entity.User
	var createdTimeStr, updatedTimeStr *string

	err := r.db.QueryRowContext(ctx, database.FindUserByIdQuery, id).Scan(
		&account.Id,
		&account.StatusOTP,
		&account.Nik,
		&account.Photo,
		&account.FirstName,
		&account.LastName,
		&account.Username,
		&account.Email,
		&account.Gender,
		&account.Address,
		&account.PhoneNumber,
		&account.EmailVerifiedAt,
		&account.Status,
		&account.CreatedBy,
		&account.UpdatedBy,
		&createdTimeStr,
		&updatedTimeStr,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	layout := "2006-01-02 15:04:05"
	if createdTimeStr != nil {
		if t, err := time.Parse(layout, *createdTimeStr); err == nil {
			account.CreatedTime = &t
		} else {
			fmt.Println("error parsing created_time:", err)
		}
	}
	if updatedTimeStr != nil {
		if t, err := time.Parse(layout, *updatedTimeStr); err == nil {
			account.UpdatedTime = &t
		} else {
			fmt.Println("error parsing updated_time:", err)
		}
	}

	return &account, nil
}

// repository/user_repository_impl.go
func (r *userRepositoryDB) FindAll(ctx context.Context, params entity.UserQuery) ([]entity.User, int, error) {
	query := "SELECT id, nik, photo, first_name, last_name, username, email, gender, address, phone_number, email_verified_at, created_by, updated_by, created_time, updated_time, status FROM users where 1=1"
	countQuery := "SELECT COUNT(*) FROM users where 1=1"
	args := []interface{}{}
	countArgs := []interface{}{}

	if params.Status != nil {
		query += " AND status = ?"
		countQuery += " AND status = ?"
		args = append(args, *params.Status)
		countArgs = append(countArgs, *params.Status)
	}

	order := "id DESC"
	if params.SortBy != "" {
		order = params.SortBy
		if params.SortOrder != "" {
			order += " " + params.SortOrder
		}
	}
	query += " ORDER BY " + order + " LIMIT ? OFFSET ?"
	args = append(args, params.Limit, params.Offset)

	// Count total
	var count int
	if err := r.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&count); err != nil {
		return nil, 0, err
	}
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		var createdTimeStr, updatedTimeStr *string

		err := rows.Scan(
			&user.Id, &user.Nik, &user.Photo, &user.FirstName, &user.LastName,
			&user.Username, &user.Email, &user.Gender, &user.Address,
			&user.PhoneNumber, &user.EmailVerifiedAt, &user.CreatedBy,
			&user.UpdatedBy, &createdTimeStr, &updatedTimeStr, &user.Status,
		)
		if err != nil {
			return nil, 0, err
		}

		// Parsing created_time dan updated_time (jika tidak nil)
		layout := time.RFC3339 // Ganti jika format di DB berbeda
		if createdTimeStr != nil {
			if t, err := time.Parse(layout, *createdTimeStr); err == nil {
				user.CreatedTime = &t
			}
		}
		if updatedTimeStr != nil {
			if t, err := time.Parse(layout, *updatedTimeStr); err == nil {
				user.UpdatedTime = &t
			}
		}

		users = append(users, user)
	}

	return users, count, nil
}

func (r *userRepositoryDB) GetUserRoles(ctx context.Context, userId int64) ([]entity.RoleUsers, []entity.RolePrivileges, error) {
	var roles []entity.RoleUsers
	var privileges []entity.RolePrivileges

	rows, err := r.db.QueryContext(ctx, database.FindRoleQuery, userId)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var roleUser entity.RoleUsers
		if err := rows.Scan(
			&roleUser.Id,
			&roleUser.UserId,
			&roleUser.RolesId.Id,
			&roleUser.RolesId.Name,
			&roleUser.RolesId.Code,
			&roleUser.RolesId.CreatedBy,
			&roleUser.RolesId.UpdatedBy,
			&roleUser.RolesId.CreatedTime,
			&roleUser.RolesId.UpdatedTime,
			&roleUser.RolesId.Status,
			&roleUser.CreatedBy,
			&roleUser.UpdatedBy,
			&roleUser.CreatedTime,
			&roleUser.UpdatedTime,
			&roleUser.Status); err != nil {
			return nil, nil, err
		}
		roles = append(roles, roleUser)
	}

	privRows, err := r.db.QueryContext(ctx, database.FindPrivilegeQuery, userId)
	if err != nil {
		return nil, nil, err
	}
	defer privRows.Close()

	for privRows.Next() {
		var privilege entity.RolePrivileges
		if err := privRows.Scan(
			&privilege.Id,
			&privilege.Role,
			&privilege.Action,
			&privilege.Uri,
			&privilege.Method,
			&privilege.CreatedBy,
			&privilege.UpdatedBy,
			&privilege.CreatedTime,
			&privilege.UpdatedTime,
			&privilege.Status); err != nil {
			return nil, nil, err
		}
		privileges = append(privileges, privilege)
	}

	return roles, privileges, nil
}

func (r *userRepositoryDB) GetOTPByCode(ctx context.Context, otp string, userId int) (*entity.UserOtps, error) {
	var detailUserOtps entity.UserOtps
	err := r.db.QueryRowContext(ctx, database.FindUserOtpsByOTP, otp, userId).Scan(
		&detailUserOtps.Id,
		&detailUserOtps.User_id,
		&detailUserOtps.Otp,
		&detailUserOtps.Expired_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &detailUserOtps, nil
}

func (r *userRepositoryDB) FindAccountByUserId(ctx context.Context, userId int) (*entity.UserRoles, error) {
	var account entity.UserRoles
	err := r.db.QueryRowContext(ctx, database.FindAccountByUserIdQuery, userId).Scan(
		&account.Id,
		&account.Nik,
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
		&account.RoleCode,
		&account.Status)
	if err != nil {
		if err == sql.ErrNoRows {
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
		&account.RoleCode,
		&account.Status)
	if err != nil {
		if err == sql.ErrNoRows {
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
		&account.Nik,
		&account.StatusOTP,
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
		&account.RoleCode,
		&account.Status)
	if err != nil {
		if err == sql.ErrNoRows {
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
		user.Nik,
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

func (r *userRepositoryDB) CreateRoleUser(ctx context.Context, userId int) error {
	roles_id := 2
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
