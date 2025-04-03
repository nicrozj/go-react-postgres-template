package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL не задан")
	}

	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("Ошибка при подключении к БД: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Не удалось проверить подключение: %v", err)
	}

	log.Println("База данных успешно подключена")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}