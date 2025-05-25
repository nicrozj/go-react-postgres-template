package repository

import (
	"backend/internal/config"
	"backend/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo() AuthRepositoryInterface {
	return &AuthRepo{
		db: config.NewAppConfig().DB,
	}
}

func (r AuthRepo) CreateUser(user *models.User) (int, error)
func (r AuthRepo) LogoutUser(userID int) (int, error)
func (r AuthRepo) DeleteUser(userID int) (int, error)
func (r AuthRepo) GetUserByID(userID int) (*models.User, int, error)
func (r AuthRepo) LoginUser(user *models.User) (*models.TokenResponse, int, error)
func (r AuthRepo) GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, int, error)
func (r AuthRepo) getAuthTokens(userID int) (*models.TokenResponse, int, error)
