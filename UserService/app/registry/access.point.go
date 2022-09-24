package registry

import (
	"SepFirst/UserService/app/interface/persistence/rdbms/concrete"
	"SepFirst/UserService/app/interface/persistence/rdbms/mysqldb"
	"SepFirst/UserService/app/usecase"
	"SepFirst/UserService/app/usecase/interactor"
	"database/sql"
)

func NewQuerier(NeedTransaction bool, db *sql.DB) mysqldb.Querier {
	conn := BuildTx(db, NeedTransaction)
	return mysqldb.Querier{DB: conn}
}

// Flag true returns the Tx
func BuildTx(db *sql.DB, Flag bool) concrete.DBTX {
	if Flag {
		tx, err := db.Begin()
		if err != nil {
			return nil
		}
		return &concrete.TxConn{DB: tx}
	}
	return &concrete.DBConn{DB: db}
}

type UserAccessPoint struct {
	Service usecase.UserService
}

func BuildUserAccessPoint(NeedTransaction bool, db *sql.DB) *UserAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewUserUsecase(&querier)

	return &UserAccessPoint{
		Service: usecaselayer,
	}
}

type AdminAccessPoint struct {
	Service usecase.AdminService
}

func BuildAdminAccessPoint(NeedTransaction bool, db *sql.DB) *AdminAccessPoint {
	querier := NewQuerier(NeedTransaction, db)
	usecaselayer := interactor.NewAdminUsecase(&querier)

	return &AdminAccessPoint{
		Service: usecaselayer,
	}
}
