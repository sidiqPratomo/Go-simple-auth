package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sidiqPratomo/DJKI-Pengaduan/database"
)

type RefreshTokenRepository interface{
	InvalidateCodes(ctx context.Context, accountId int64) error
	PostOneCode(ctx context.Context, accountId int64, code string) error
	FindOneCode(ctx context.Context, refreshToken string) (string, error)
}

type refreshTokenRepositoryDB struct{
	db DBTX
}

func NewRefreshTokenRepositoryDB(db *sql.DB) refreshTokenRepositoryDB {
	return refreshTokenRepositoryDB{
		db: db,
	}
}

func (r *refreshTokenRepositoryDB) InvalidateCodes(ctx context.Context, accountId int64) error {
	_, err := r.db.ExecContext(ctx, database.InvalidateRefreshTokensQuery, accountId)
	if err != nil {
		return err
	}

	return nil
}

func (r *refreshTokenRepositoryDB) PostOneCode(ctx context.Context, accountId int64, code string) error {
	id := uuid.New().String()
	_, err := r.db.ExecContext(ctx, database.PostOneRefreshTokenQuery, id,  accountId, code)
	if err != nil {
		return err
	}

	return nil
}

func (r *refreshTokenRepositoryDB) FindOneCode(ctx context.Context, refreshToken string) (string, error) {
	err := r.db.QueryRowContext(ctx, database.FindOneRefreshTokenQuery, refreshToken).Scan(&refreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return refreshToken, nil
}