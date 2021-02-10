package app

import (
	"github.com/gin-gonic/gin"
	"gotdd/pkg/api"
	"log"
	"net/http"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		context.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Content-Type", "application/json")
		var newUser api.NewUserRequest

		err := context.ShouldBindJSON(&newUser)
		if err != nil {
			log.Printf("handle error: %v", err)
			context.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.userService.New(newUser)
		if err != nil {
			log.Printf("service error: %v", err)
			context.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "status",
			"data":   "new user created",
		}

		context.JSON(http.StatusOK, response)

	}
}
