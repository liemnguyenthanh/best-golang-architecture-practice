package server

import (
	"api-instagram/config"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	cfg   *config.Configuration
	db    *gorm.DB
	gin   *gin.Engine
}

func NewServer(cfg *config.Configuration, db *gorm.DB) *Server {
	r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

	return &Server{cfg: cfg, db: db, gin: r}
}

func (s *Server) Run() *Server {
	fmt.Println("Server is runing at", s.cfg.Port)

	s.gin.Run(":" + s.cfg.Port)

	return s
}
