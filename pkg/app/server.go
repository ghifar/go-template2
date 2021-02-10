package app

import (
	"github.com/gin-gonic/gin"
	"gotdd/pkg/api"
	"log"
)

type Server struct {
	router        *gin.Engine
	userService   api.UserService
	weightService api.WeightService
}

func NewServer(
	router *gin.Engine,
	userService api.UserService,
	weightService api.WeightService,
) *Server {
	return &Server{
		router:        router,
		userService:   userService,
		weightService: weightService,
	}
}

func (s *Server) Run() error {
	//run initialize routes func
	r := s.Routes()

	//run server router
	err := r.Run()

	if err != nil {
		log.Printf("Server error calling router, %v", err)
		return err
	}

	return nil
}
