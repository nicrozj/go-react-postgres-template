package main

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.GET("/greet", handlers.Greet)
	r.Run(":8080")
}
