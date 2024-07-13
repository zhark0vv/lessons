package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"lessons/cleanarch/internal/domain"
)

type adp interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// UserRepository - репозиторий для работы с пользователями в SQLite
type UserRepository struct {
	adp adp
}

func NewSQLiteUserRepository(adp adp) *UserRepository {
	return &UserRepository{
		adp: adp,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := r.adp.ExecContext(ctx, `INSERT OR REPLACE INTO users (id, name, email, password) VALUES (?, ?, ?, ?)`,
		user.ID.String(), user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row := r.adp.QueryRowContext(ctx, `SELECT id, name, email, password FROM users WHERE id = ?`, id.String())

	var user domain.User
	var userID string
	err := row.Scan(&userID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	user.ID, err = uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
