package main

import (
	"net/http"

	"github.com/ekholme/dinner-helper/srvr"
	"github.com/gin-gonic/gin"
)

func main() {
	r := srvr.NewServer()
	//h := NewHandler()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8080")

}
