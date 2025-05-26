package services

import (
	"backend/internal/models"
)

type AuthServiceInterface interface {
	CreateUser(body *models.AuthRequest) (*models.Response, *models.ErrorResponse)
	DeleteUser(userID int) (*models.Response, *models.ErrorResponse)
	LogoutUser(userID int) (*models.Response, *models.ErrorResponse)
	GetUserByID(userID int) (*models.Response, *models.ErrorResponse)
	LoginUser(body *models.AuthRequest) (*models.TokenResponse, *models.ErrorResponse)
	GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, *models.ErrorResponse)
}
