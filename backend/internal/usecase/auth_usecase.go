package usecase

import (
	"backend/internal/entity"
	"backend/internal/pkg/jwt"
	"backend/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepo  repository.AuthRepository
	TokenRepo repository.TokenRepository
}

func NewAuthUsecase(AuthRepo repository.AuthRepository, TokenRepo repository.TokenRepository) *AuthUsecase {
	return &AuthUsecase{
		AuthRepo:  AuthRepo,
		TokenRepo: TokenRepo,
	}
}

func (uc *AuthUsecase) Login(email, password string) (*entity.Token, error) {
	user, err := uc.AuthRepo.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("email is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, fmt.Errorf("password is incorrect")
	}

	accessToken, refreshToken, refreshExp, err := jwt.GenerateTokens(user.ID)
	if err != nil {
		return nil, err
	}

	err = uc.TokenRepo.SaveRefreshToken(refreshToken, refreshExp, user)
	if err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}

	return &entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *AuthUsecase) Registration(email, password string) (*entity.Token, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    email,
		Password: string(hashPassword),
	}

	if err := uc.AuthRepo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("the user already exists")
	}

	accessToken, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &entity.Token{AccessToken: accessToken}, nil
}
