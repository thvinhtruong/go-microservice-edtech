package db

import (
	"context"
	"database/sql"
	"log"
)

type Transactioner interface {
	Rollback() error
	Commit() error
	EnableTx(txFunc func() error) error
}

type TxStore struct {
	*Queries
	db *sql.DB
	Transactioner
}

func NewTxStore(db *sql.DB) *TxStore {
	return &TxStore{
		Queries: New(db),
		db:      db,
	}
}

func (t *TxStore) Rollback() error {
	return t.Queries.db.(*sql.Tx).Rollback()
}

func (t *TxStore) Commit() error {
	return t.Queries.db.(*sql.Tx).Commit()
}

// a context and a callback function as input, start a transaction
// pass the transaction to the callback function.
// if the callback function returns an error, rollback the transaction
// returns a transaction object or an error
func (t *TxStore) EnableTx(txFunc func() error) error {
	tx, err := t.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	t.Queries = New(tx)
	defer func() {
		if p := recover(); p != nil {
			log.Println("found problem and rollback:", p)
			tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Println("found problem and rollback:", err)
			tx.Rollback()
		} else {
			log.Println("nothing happens, commit")
			err = tx.Commit()
		}
	}()

	err = txFunc()
	return err
}
