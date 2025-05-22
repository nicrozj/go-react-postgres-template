package repository

import (
	"backend/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) GetUserByEmail(email string) (*entity.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	user := &entity.User{}
	err := r.db.Get(user, query, email)
	return user, err
}

func (r *AuthRepo) CreateUser(user *entity.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
}

func (r *AuthRepo) SaveRefreshToken(refreshToken string, refreshExp time.Time, user *entity.User) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.ID, refreshToken, refreshExp)
	return err
}
