package jwt

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func getExpiration() time.Duration {
	expStr := os.Getenv("JWT_EXPIRATION")
	if expStr == "" {
		return time.Hour * 24 // Default to 24 hours
	}
	
	exp, err := strconv.Atoi(expStr[:len(expStr)-1])
	if err != nil {
		return time.Hour * 24
	}
	
	unit := expStr[len(expStr)-1:]
	switch unit {
	case "h":
		return time.Hour * time.Duration(exp)
	case "d":
		return time.Hour * 24 * time.Duration(exp)
	default:
		return time.Hour * 24
	}
}

func GenerateToken(userID int) (string, error) {
	expiration := getExpiration()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiration).Unix(),
	})

	return token.SignedString(getSecretKey())
}

func GenerateTokens(userID int) (accessToken, refreshToken string, expiresAt time.Time, err error) {
	expiration := getExpiration()
	expiresAt = time.Now().Add(expiration)
	
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expiresAt.Unix(),
	}).SignedString(getSecretKey())

	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshExp := time.Now().Add(7 * 24 * time.Hour)
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     refreshExp.Unix(),
	}).SignedString(getSecretKey())

	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, refreshExp, err
}

func RefreshAccessToken(refreshToken string) (newAccessToken string, err error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}
	
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", errors.New("invalid user_id in token")
	}
	
	return GenerateToken(int(userID))
}

func ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return getSecretKey(), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token")
	}

	return int(userID), nil
}
