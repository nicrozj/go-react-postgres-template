package config

import "os"

type JWTConfig struct {
	SecretKet string
}

func NewJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKet: os.Getenv("JWT_SECRET_KEY"),
	}
}
