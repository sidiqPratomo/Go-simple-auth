package repository

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Transaction interface {
	Rollback() error
	Commit() error
	BeginTx() (Transaction, error)
	UserRepository() UserRepository
	RefreshTokenRepository() RefreshTokenRepository
}

type SqlTransaction struct {
	db *sql.DB
	tx *sql.Tx
}

func NewSqlTransaction(db *sql.DB) *SqlTransaction {
	return &SqlTransaction{
		db: db,
	}
}

func (s *SqlTransaction) BeginTx() (Transaction, error) {
	tx, err := s.db.Begin()
	return &SqlTransaction{db: s.db, tx: tx}, err
}

func (s *SqlTransaction) Rollback() error {
	return s.tx.Rollback()
}

func (s *SqlTransaction) Commit() error {
	return s.tx.Commit()
}

func (s *SqlTransaction) UserRepository() UserRepository {
	return &userRepositoryDB{
		db: s.tx,
	}
}

func (s *SqlTransaction) RefreshTokenRepository() RefreshTokenRepository {
	return &refreshTokenRepositoryDB{
		db: s.tx,
	}
}