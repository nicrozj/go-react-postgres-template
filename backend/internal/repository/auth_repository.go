package repository

import (
	"backend/internal/entity"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}
