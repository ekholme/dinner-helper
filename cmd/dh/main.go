package main

import (
	"github.com/ekholme/dinner-helper/frstr"
	"github.com/ekholme/dinner-helper/srvr"
	"github.com/gin-gonic/gin"
)

func main() {
	ms := frstr.NewMealService()
	mh := srvr.NewMealHandler(ms)
	r := gin.Default()

	r.GET("/meal", mh.GetAllMeals)
	r.POST("/meal", mh.CreateMeal)

	r.Run(":8080")

}
