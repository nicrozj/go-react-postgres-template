package main

import (
	"backend/internal/delivery"
	"backend/internal/repository"
	"backend/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "postgres"
	dbPort     = "5432"
	dbUser     = "admin"
	dbPassword = "password"
	dbName     = "db"
)

func main() {
	db, err := sqlx.Connect("postgres",
		"host="+dbHost+" port="+dbPort+" user="+dbUser+" password="+dbPassword+" dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authRepo := repository.NewAuthRepo(db)
	tokenRepo := repository.NewTokenRepo(db)
	authUC := usecase.NewAuthUsecase(authRepo, tokenRepo)
	authHandler := delivery.NewAuthHandler(*authUC)

	r := gin.Default()

	r.POST("/registration", authHandler.Registration)
	r.POST("/login", authHandler.Login)

	r.Run(":8080")
}
