package main

import (
	"backend/internal/api"
	"backend/internal/config"

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

	server := api.NewServer()
	if err := server.Engine.Run(":8080"); err != nil {
		panic(err)
	}
}
