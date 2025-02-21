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
	engine := gin.Default()
	engine.Use(cors.Default())
	return &Server{
		engine: engine,
		cfg:    config.GetConfig(),
	}
}

func (s Server) Run() {
	_ = s.engine.SetTrustedProxies(nil)
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", s.cfg.HttpPort)
	}
	if environment := os.Getenv("ENVIRONMENT"); environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	s.setupRoutes()

	if err := s.engine.Run(":" + port); err != nil {
		log.Fatalf("Running HTTP server err: %v", err)
	}
}

func (s Server) setupRoutes() {
	v1 := s.engine.Group("/api/v1")
	taskHttp.Routes(v1)
}
