package inmemory

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"lessons/cleanarch/internal/domain"
)

// 4. Repository - репозиторий.
// Репозиторий - это объект, который инкапсулирует логику доступа к данным
// и предоставляет интерфейс для работы с этими данными

type UserRepository struct {
	users map[uuid.UUID]*domain.User
}

func NewInMemoryUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]*domain.User),
	}
}

func (r *UserRepository) Save(_ context.Context, user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(_ context.Context, id uuid.UUID) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
