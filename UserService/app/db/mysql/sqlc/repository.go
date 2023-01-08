package db

import (
	"context"
	"database/sql"
)

type Repository interface {
	RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserResult, error)
	LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserResult, error)
	RegisterTutor(ctx context.Context, arg RegisterTutorParams) (RegisterTutorResult, error)
	LoginTutor(ctx context.Context, arg LoginTutorParams) (LoginTutorResult, error)
	FilterTutors(ctx context.Context, arg FindTutorsMatchParams) ([]int32, error)
}

func NewRepository(db *sql.DB) Repository {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}
