package repository

import (
	"backend/internal/entity"
	"time"

	"github.com/jmoiron/sqlx"
)

type TokenRepository interface {
	SaveRefreshToken(refreshToken string, refershExp time.Time, user *entity.User) error
}

type TokenRepo struct {
	db *sqlx.DB
}

func NewTokenRepo(db *sqlx.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (r *TokenRepo) SaveRefreshToken(refreshToken string, refreshExp time.Time, user *entity.User) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.ID, refreshToken, refreshExp)
	return err
}
