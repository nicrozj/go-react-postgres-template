package config

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewDBClient() *sqlx.DB {

	db, err := sqlx.Connect("postgres", Envs.DB_URL)
	if err != nil {
		panic(err)
	}

	return db
}
