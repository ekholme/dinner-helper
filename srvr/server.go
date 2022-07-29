package srvr

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//server struct that holds the pieces we're using
type Server struct {
	router *gin.Engine
	mh     MealHandler
}

//function to create a new server
func NewServer(router *gin.Engine, mh MealHandler) *Server {

	return &Server{
		router: router,
		mh:     mh,
	}

}

//handle index
//this may not be the best place to do this
func (s *Server) handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}

//method to run the server
func (s *Server) Run() {
	//register index
	s.router.GET("/", s.handleIndex)

	s.registerMealRoutes()

	s.router.Run(":8080")
}
