package main

import (
	"github.com/ekholme/dinner-helper/frstr"
	"github.com/ekholme/dinner-helper/srvr"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	r.LoadHTMLGlob("srvr/html/templates/*.html")
}

func main() {
	ms := frstr.NewMealService()
	mh := srvr.NewMealHandler(ms)
	s := srvr.NewServer(r, mh)

	// type server struct {
	// 	router *gin.Engine
	// 	mh     srvr.MealHandler
	// }

	// var s = &server{
	// 	router: r,
	// 	mh:     mh,
	// }

	// s.router.GET("/meal", s.mh.GetAllMeals)

	// r.GET("/meal", mh.GetAllMeals)
	// r.POST("/meal", mh.CreateMeal)

	// r.Run(":8080")

	s.Run()

}
