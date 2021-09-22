package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tmdb/pkg/api"
)

type Server struct {
	router  *gin.Engine
	service *api.Service
}

func NewServer(router *gin.Engine, service *api.Service) *Server {
	return &Server{
		router:  router,
		service: service,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
