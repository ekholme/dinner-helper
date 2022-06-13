package srvr

import (
	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	//the server instance to create, using gin
	router *gin.Engine

	//services used by routes
	//just mealservice for now
	MealService dh.MealService
}

//creates a new server
func NewHandler() *Handler {
	return &Handler{}
}

func NewServer() *gin.Engine {
	return gin.Default()
}
