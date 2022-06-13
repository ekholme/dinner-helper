package http

import (
	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

type Server struct {
	//the server instance to create, using gin
	router *gin.Engine

	//services used by routes
	//just mealservice for now
	MealService dh.MealService
}

func NewServer() *Server {
	s := &Server{
		router: gin.Default(),
	}

	return s
}
