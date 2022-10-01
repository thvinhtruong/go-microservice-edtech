package db

import (
	"context"
	"database/sql"
)

type TxStore struct {
	*Queries
	db *sql.DB
}

func NewTxStore(db *sql.DB) *TxStore {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}

// a context and a callback function as input, start a transaction
// pass the transaction to the callback function.
// if the callback function returns an error, rollback the transaction
// returns a transaction object or an error
func (t *TxStore) EnableTx(ctx context.Context, fn func(*Queries) error) error {
	// using nil for the default isolation level
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}
