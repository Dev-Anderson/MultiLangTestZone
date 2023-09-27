package server

import (
	"aut-jwt/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8099",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.HandlerRoutes(s.server)
	log.Fatal(router.Run(":" + s.port))
}
