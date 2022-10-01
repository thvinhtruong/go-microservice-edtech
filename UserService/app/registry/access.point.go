package registry

import (
	db "SepFirst/UserService/app/interface/db/mysql/sqlc"
	"SepFirst/UserService/app/usecase"
	"SepFirst/UserService/app/usecase/interactor"
	"database/sql"
)

func NewQueries(database *sql.DB) *db.Queries {
	return db.New(database)
}

type UserAccessPoint struct {
	Service usecase.UserService
}

func BuildUserAccessPoint(db *sql.DB) *UserAccessPoint {
	queries := NewQueries(db)
	usecaselayer := interactor.NewUserUsecase(&queries)

	return &UserAccessPoint{
		Service: usecaselayer,
	}
}
