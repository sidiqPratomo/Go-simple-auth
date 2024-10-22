package repository

import (
	"context"
	"database/sql"

	"github.com/sidiqPratomo/DJKI-Pengaduan/database"
	"github.com/sidiqPratomo/DJKI-Pengaduan/entity"
)

type UserRepository interface {
	PostOneUser(ctx context.Context, account entity.User) (*int, error)
	GetAllUser(ctx context.Context, user entity.User) ([]entity.User, error)
}

type userRepositoryDB struct {
	db DBTX
}

func NewUserRepositoryDB(db *sql.DB) userRepositoryDB {
	return userRepositoryDB{
		db: db,
	}
}

func (r *userRepositoryDB) PostOneUser(ctx context.Context, user entity.User) (*int, error) {
	var userId int

	err := r.db.QueryRowContext(ctx, database.PostOneAccountQuery, 
		user.Email, 
		user.Password, 
		user.FirstName, 
		user.LastName, 
		user.PhoneNumber, 
		user.CreatedBy,
	).Scan(&userId)
	if err != nil {
		return nil, err
	}

	return &userId, nil
}

func (r *userRepositoryDB) GetAllUser(ctx context.Context, user entity.User) ([]entity.User, error){
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
		if err != nil{
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}