package usecase

import "backend/internal/repository"

type TokenUsecase struct {
	AuthRepo  repository.AuthRepository
	TokenRepo repository.TokenRepository
}

func NewTokenUsecase(TokenRepo repository.TokenRepository) *TokenUsecase {
	return &TokenUsecase{
		TokenRepo: TokenRepo,
	}
}
