package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	taskHttp "todo-list/internal/task/http"
	"todo-list/pkg/config"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
		cfg:    config.GetConfig(),
	}
}

func (s Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", s.cfg.HttpPort)
	}

	s.engine.Use(corsMiddleware())
	s.setupRoutes()

	if err := s.engine.Run(":" + port); err != nil {
		log.Fatalf("Running HTTP server err: %v", err)
	}
}

func (s Server) setupRoutes() {
	v1 := s.engine.Group("/api/v1")
	taskHttp.Routes(v1)
}

func corsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}                                                // Allow requests from any origin
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"} // Allow specific HTTP methods
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}                    // Allow specific headers
	return cors.New(corsConfig)
}
