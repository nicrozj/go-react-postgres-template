package main

import (
	"backend/internal/config"
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	_, err = config.ParseEnvs()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/greet", handlers.Greet)
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
