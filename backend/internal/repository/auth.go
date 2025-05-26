package repository

import (
	"backend/internal/config"
	"backend/internal/models"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepo struct {
	db  *sqlx.DB
	jwt config.JWTConfig
}

func NewAuthRepo() AuthRepositoryInterface {
	return &AuthRepo{
		db:  config.NewAppConfig().DB,
		jwt: config.NewJWTConfig(),
	}
}

func (r AuthRepo) CreateUser(user *models.User) (int, error) {
	var row int
	query := "SELECT count(username) FROM users WHERE username = $1"
	err := r.db.QueryRow(query, user.Username).Scan(&row)
	if err != nil && err != sql.ErrNoRows {
		return http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	if row != 0 {
		return http.StatusBadRequest, fmt.Errorf("username already taken, please use different username")
	}

	query = "INSERT INTO users(username, password_hash) VALUES($1, $2) RETURNING id"
	hashPassword, err := getHashPassword(user.PasswordHash)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	_, err = r.db.Exec(query, user.Username, string(hashPassword))
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	return http.StatusOK, nil
}

func (r AuthRepo) GetUserByID(userID int) (*models.User, int, error) {
	var user models.User
	query := "SELECT id, username, password_hash FROM users WHERE id = $1"
	err := r.db.QueryRow(query, userID).Scan(&user.Id, &user.Username, &user.PasswordHash)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	return &user, http.StatusOK, nil
}

func (r AuthRepo) DeleteUser(userID int) (int, error) {
	query := "DELETE FROM users WHERE id = $1"

	_, err := r.db.Exec(query, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}

func (r AuthRepo) LoginUser(user *models.User) (*models.TokenResponse, int, error) {
	existUser := &models.User{}
	query := "SELECT id, username, password_hash FROM users WHERE username = $1"
	err := r.db.QueryRow(query, user.Username).Scan(&existUser.Id, &existUser.Username, &existUser.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusBadRequest, fmt.Errorf("please check credentials")
		}
		return nil, http.StatusBadRequest, err
	}

	isValid := checkPassword(existUser.PasswordHash, user.PasswordHash)
	if !isValid {
		return nil, http.StatusUnauthorized, fmt.Errorf("incorrect password, please try again")
	}

	return r.getAuthTokens(existUser.Id)
}

func (r AuthRepo) LogoutUser(userID int) (int, error) {
	query := "DELETE FROM refresh_tokens WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("please try again later")
	}
	return http.StatusOK, nil
}

func (r AuthRepo) GenerateTokens(userID int, oldRefreshToken string) (*models.TokenResponse, int, error) {
	var dbRefreshToken string

	query := "SELECT refresh_token FROM refresh_tokens WHERE user_id = $1"

	err := r.db.QueryRow(query, userID).Scan(&dbRefreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusUnauthorized, fmt.Errorf("please login again")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("please login again")
	}

	if dbRefreshToken != oldRefreshToken {
		return nil, http.StatusUnauthorized, fmt.Errorf("invalid token, please login again")
	}

	return r.getAuthTokens(userID)
}

func (r AuthRepo) getAuthTokens(userID int) (*models.TokenResponse, int, error) {
	accessTokenExpire := time.Now().Add(time.Duration(config.Envs.HTTP_ACCESS_TOKEN_EXPIRE) * time.Minute).Unix()
	accessToken, err := getToken(userID, accessTokenExpire)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	refreshTokenExpire := time.Now().Add(time.Duration(config.Envs.HTTP_REFRESH_TOKEN_EXPIRE) * time.Minute).Unix()
	refreshToken, err := getToken(userID, refreshTokenExpire)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	query := `
    INSERT INTO refresh_tokens (user_id, refresh_token, expire_time, created_at)
    VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
    ON CONFLICT (user_id)
    DO UPDATE SET
        refresh_token = EXCLUDED.refresh_token,
        expire_time = EXCLUDED.expire_time,
        created_at = CURRENT_TIMESTAMP
`
	_, err = r.db.Exec(query, userID, refreshToken, refreshTokenExpire)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("please try again later")
	}

	return &models.TokenResponse{AccessToken: accessToken, RefreshToken: refreshToken}, http.StatusOK, err
}

func getToken(userID int, expireTime int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expireTime,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Envs.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, nil
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
