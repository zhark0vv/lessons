package domain

import (
	"time"

	"github.com/google/uuid"
)

// 1. Entity - сущность.
// Сущность имеет уникальный идентификатор и жизненный цикл

type User struct {
	ID       uuid.UUID
	Address  Address
	Name     string
	Email    string
	Password string
}

func (u *User) ChangeEmail(newEmail string) {
	u.Email = newEmail
}

// 5. Domain Event - событие домена.
// Событие домена - это событие, которое произошло в предметной области

type UserRegistered struct {
	UserID     uuid.UUID
	OccurredAt time.Time
}

func NewUserRegistered(userID uuid.UUID) UserRegistered {
	return UserRegistered{
		UserID:     userID,
		OccurredAt: time.Now(),
	}
}
