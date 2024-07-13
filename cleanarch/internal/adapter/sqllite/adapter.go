package sqllite

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Adapter struct {
	db *sql.DB
}

func NewSQLLiteAdapter(dataSourceName string) (*Adapter, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return a.db.ExecContext(ctx, query, args...)
}

func (a *Adapter) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return a.db.QueryRowContext(ctx, query, args...)
}
