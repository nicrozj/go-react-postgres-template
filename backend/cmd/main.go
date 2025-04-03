package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type User struct {
  ID   int    `db:"id"`
  Name string `db:"name"`
}

func main() {
  err := godotenv.Load()
  if err != nil {
      log.Println("Не удалось загрузить .env файл:", err)
  }

  router := gin.Default()

  db.InitDB()
	defer db.CloseDB()

	var users []User
	err = db.DB.Select(&users, "SELECT * FROM users")

  for _, user := range users {
    log.Printf("User: %s, ID: %d ", user.Name, user.ID)
  }

	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}

  handlers.SetupRoutes(router)

  router.Run()
}