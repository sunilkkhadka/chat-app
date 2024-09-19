package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sunilkkhadka/chat-app/internal/config"
)

type Server struct {
	Cfg *config.Config
	Gin *gin.Engine
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Cfg: cfg,
		Gin: gin.Default(),
	}
}

func (server Server) Run(address string) error {
	return server.Gin.Run(":" + address)
}
