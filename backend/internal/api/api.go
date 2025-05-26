package api

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	s := &Server{
		Engine: gin.Default(),
	}
	s.mountMiddlewares()
	s.mountHandlers()
	return s
}

func (s *Server) mountMiddlewares() {
	s.Engine.Use(func(c *gin.Context) {
		c.Set("timeout", 1*time.Minute)
		c.Next()
	})

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{config.Envs.WEB_URL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	s.Engine.Use(ginCorsMiddleware(corsMiddleware))

	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func ginCorsMiddleware(corsMiddleware *cors.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	}
}

func (s *Server) mountHandlers() {
	authHandlers := handlers.NewAuthHandlers()

	authGroup := s.Engine.Group("/auth")
	{
		authGroup.GET("/greet", authHandlers.Greet)
		authGroup.POST("/login", authHandlers.LoginUser)
		authGroup.POST("/registration", authHandlers.CreateUser)
	}

	protectedGroup := s.Engine.Group("/")
	protectedGroup.Use(jwtMiddleware())
	{
		protectedGroup.POST("/auth/logout", authHandlers.LogoutUser)
		protectedGroup.POST("/tokens/refresh", authHandlers.RefreshToken)
		protectedGroup.DELETE("/users", authHandlers.DeleteUser)
		protectedGroup.GET("/users/me", authHandlers.GetUserByID)
	}
}
