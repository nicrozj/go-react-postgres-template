package main

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbConfig := config.LoadDBConfig()
	db, err := sqlx.Connect(
		"postgres",
		"host="+dbConfig.Host+" port="+dbConfig.Port+
			" user="+dbConfig.User+" password="+dbConfig.Password+
			" dbname="+dbConfig.Name+" sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/greet", handlers.Greet)
	r.Run(":8080")
}
