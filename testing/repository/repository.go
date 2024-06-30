package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{conn: conn}
}

func (r *Repository) GetData(ctx context.Context, id int) (string, error) {
	var data string
	err := r.conn.QueryRow(ctx, "SELECT data FROM test_table WHERE id=$1", id).Scan(&data)
	if err != nil {
		return "", err
	}
	return data, nil
}
