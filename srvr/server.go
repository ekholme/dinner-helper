package srvr

import "github.com/gin-gonic/gin"

//TODO
//create a server that holds handlers, routes, etc
//see https://github.com/benbjohnson/wtf/blob/main/http/server.go

//not sure why it's not recognizing the gin engine
//see below
//https://levelup.gitconnected.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189
//try creating a Run() method that does a lot of the stuff I want to do when running
//see the example above
type Server struct {
	router *gin.Engine
	mh     MealHandler
}

//for some reason this isn't currently working for me
func NewServer(router *gin.Engine, mh MealHandler) *Server {

	return &Server{
		router: router,
		mh:     mh,
	}

}
