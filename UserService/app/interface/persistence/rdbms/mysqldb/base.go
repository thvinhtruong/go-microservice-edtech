package mysqldb

import (
	"SepFirst/UserService/app/interface/persistence/rdbms/concrete"
)

type Querier struct {
	DB concrete.DBTX
}

func NewQuerier(db concrete.DBTX) *Querier {
	return &Querier{
		DB: db,
	}
}

func (q *Querier) EnableTx(txFunc func() error) error {
	return q.DB.TxEnd(txFunc)
}
