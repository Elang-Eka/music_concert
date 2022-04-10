package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// A Tx belong to the inteface layer.
type CommandTx interface {
	Commit() error
	Rollback() error
	NamedExec(name string, query string, args interface{}) (sql.Result, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

type commandTx struct {
	ctx context.Context
	tx  *sqlx.Tx
}

func initTx(ctx context.Context, tx *sqlx.Tx) CommandTx {
	return &commandTx{
		ctx: ctx,
		tx:  tx,
	}
}

func (x *commandTx) Commit() error {
	return x.tx.Commit()
}

func (x *commandTx) Rollback() error {
	return x.tx.Rollback()
}

func (x *commandTx) NamedExec(name string, query string, args interface{}) (sql.Result, error) {
	return x.tx.NamedExecContext(x.ctx, query, args)
}

func (x *commandTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return x.tx.ExecContext(x.ctx, query, args...)
}
