package usecase

import (
	"backend/internal/entity"
	"backend/internal/pkg/jwt"
	"backend/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	repo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) Registartion(email, password string) (*entity.Token, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    email,
		Password: string(hashPassword),
	}

	if err := uc.repo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("email is already use")
	}

	accessToken, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.Token{AccessToken: accessToken}, nil
}

func (uc *AuthUsecase) Login(email, password string) (*entity.Token, error) {
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	accessToken, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.Token{AccessToken: accessToken}, nil
}
