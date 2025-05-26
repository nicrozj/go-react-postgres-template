package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type AuthService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService() AuthServiceInterface {
	return &AuthService{
		repo: repository.NewAuthRepo(),
	}
}

func (s AuthService) CreateUser(body *models.AuthRequest) (*models.Response, *models.ErrorResponse) {
	user := &models.User{Username: body.Username, PasswordHash: body.Password}
	status, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status}, nil
}

func (s AuthService) DeleteUser(userId int) (*models.Response, *models.ErrorResponse) {
	status, err := s.repo.DeleteUser(userId)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status}, nil
}

func (s AuthService) LogoutUser(userID int) (*models.Response, *models.ErrorResponse) {
	status, err := s.repo.LogoutUser(userID)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status}, nil
}

func (s AuthService) GetUserByID(userID int) (*models.Response, *models.ErrorResponse) {
	user, status, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return &models.Response{Success: true, Status: status, Data: []*models.User{user}}, nil
}

func (s AuthService) LoginUser(body *models.AuthRequest) (*models.TokenResponse, *models.ErrorResponse) {
	user := models.User{Username: body.Username, PasswordHash: body.Password}
	tokenRes, status, err := s.repo.LoginUser(&user)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return tokenRes, nil
}

func (s AuthService) GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, *models.ErrorResponse) {
	tokenRes, status, err := s.repo.GenerateTokens(userID, oldRefreshToken)
	if err != nil {
		return nil, models.NewErrorResponse(status, err)
	}
	return tokenRes, nil
}
