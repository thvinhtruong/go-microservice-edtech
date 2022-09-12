package concrete

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Transactioner
}

type Transactioner interface {
	Commit() error
	Rollback() error
	TxEnd(func() error) error
	// TxEnd commits a transaction if no errors, otherwise rollback
	// txFunc is the operations wrapped in a transaction
}
