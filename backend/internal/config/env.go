package config

import (
	"fmt"
	"os"
	"strconv"
)

type AppEnvs struct {
	JWT_SECRET_KEY            string
	DB_URL                    string
	WEB_URL                   string
	HTTP_ACCESS_TOKEN_EXPIRE  int
	HTTP_REFRESH_TOKEN_EXPIRE int
}

var Envs *AppEnvs

func ParseEnvs() (*AppEnvs, error) {
	var err error

	Envs = &AppEnvs{
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
		WEB_URL:        os.Getenv("WEB_URL"),
		DB_URL:         os.Getenv("DB_URL"),
	}

	if Envs.DB_URL == "" || Envs.JWT_SECRET_KEY == "" || Envs.WEB_URL == "" {
		err = fmt.Errorf("invalid env variables in .env file")
	}

	httpRefreshTokenExpire, parseErr := strconv.Atoi(os.Getenv("HTTP_REFRESH_TOKEN_EXPIRE"))
	if parseErr != nil {
		err = parseErr
	}
	Envs.HTTP_REFRESH_TOKEN_EXPIRE = httpRefreshTokenExpire

	httpAccessTokenExpire, parseErr := strconv.Atoi(os.Getenv("HTTP_ACCESS_TOKEN_EXPIRE"))
	if parseErr != nil {
		err = parseErr
	}
	Envs.HTTP_ACCESS_TOKEN_EXPIRE = httpAccessTokenExpire

	if err != nil {
		return nil, err
	}

	return Envs, nil
}
