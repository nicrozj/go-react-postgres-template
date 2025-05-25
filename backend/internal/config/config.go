package config

import (
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	DB *sqlx.DB
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		DB: NewDBClient(),
	}
}
