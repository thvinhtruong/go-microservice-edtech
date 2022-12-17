package registry

import (
	"database/sql"
	db "server/UserService/app/db/mysql/sqlc"
	"server/UserService/app/usecase"
	"server/UserService/app/usecase/interactor"
)

func NewQueries(database *sql.DB) *db.TxStore {
	return db.NewTxStore(database)
}

type UserAccessPoint struct {
	Service usecase.UserService
}

func BuildUserAccessPoint(db *sql.DB) *UserAccessPoint {
	queries := NewQueries(db)
	usecaselayer := interactor.NewUserUsecase(queries)

	return &UserAccessPoint{
		Service: usecaselayer,
	}
}
