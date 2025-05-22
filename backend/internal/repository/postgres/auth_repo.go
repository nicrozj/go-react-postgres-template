package postgres

import (
	"backend/internal/entity"

	"github.com/jmoiron/sqlx"
)

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
