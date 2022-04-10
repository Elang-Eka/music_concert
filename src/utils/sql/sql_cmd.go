package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// A SQLHandler belong to the inteface layer.
type Command interface {
	Begin(ctx context.Context) (CommandTx, error)
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Row, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExec(ctx context.Context, name string, query string, args interface{}) (sql.Result, error)
	Close() error
}

type command struct {
	db *sqlx.DB
}

func initCommand(db *sqlx.DB) Command {
	return &command{
		db: db,
	}
}

func (c *command) Begin(ctx context.Context) (CommandTx, error) {
	tx, err := c.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return initTx(ctx, tx), nil
}

func (c *command) Close() error {
	return c.db.Close()
}

func (c *command) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

func (c *command) NamedExec(ctx context.Context, name string, query string, args interface{}) (sql.Result, error) {
	return c.db.NamedExecContext(ctx, query, args)
}

func (c *command) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return c.db.QueryxContext(ctx, query, args...)
}

func (c *command) QueryRow(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Row, error) {
	return c.db.QueryRowxContext(ctx, query, args...), nil
}
