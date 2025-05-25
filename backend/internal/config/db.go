package config

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDBConfig() DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	return dbConfig
}

func NewDBClient() *sqlx.DB {
	dbConfig := LoadDBConfig()

	db, err := sqlx.Connect(
		"postgres",
		"host="+dbConfig.Host+" port="+dbConfig.Port+
			" user="+dbConfig.User+" password="+dbConfig.Password+
			" dbname="+dbConfig.Name+" sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
