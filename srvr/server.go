package srvr

import "github.com/gin-gonic/gin"

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

//method to run the server
func (s *Server) Run() {
	s.router.GET("/meal", s.mh.GetAllMeals)
	s.router.POST("/meal", s.mh.CreateMeal)

	s.router.Run(":8080")
}
