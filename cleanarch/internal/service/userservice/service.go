package service

import (
	"context"

	"github.com/google/uuid"
	"lessons/cleanarch/internal/domain"
)

type cacheProvider interface {
	Get(ctx context.Context, key string) (string, error)
}

type repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
}

type UserService struct {
	userRepository repository
	cacheProvider  cacheProvider
}

func NewUserService(userRepo repository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (s *UserService) ChangePassword(ctx context.Context, userID uuid.UUID, newPassword string) error {
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	user.Password = newPassword
	return s.userRepository.Save(ctx, user)
}
