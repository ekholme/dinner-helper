package main

import (
	"net/http"

	"github.com/ekholme/dinner-helper/frstr"
	"github.com/ekholme/dinner-helper/srvr"
	"github.com/gin-gonic/gin"
)

func main() {
	r := srvr.NewServer()
	//find a way to reorgnize this code so main isn't dependent on frstr
	//see ben johnson stuff
	ms := frstr.NewMealService()
	h := srvr.NewMealHandler(ms)

	//see notes in srvr/meal.go for how to revise these functions
	//this isn't currently working
	r.POST("/meal", h.CreateMeal)

	r.GET("/meal", h.FindAllMeals)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.Run(":8000")

}
