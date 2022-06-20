package srvr

import (
	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

//might need to make a handler interface that implements the methods i want?
//see daren controller for examples
type MealHandler interface {
	CreateMeal(c *gin.Context)
	FindAllMeals(c *gin.Context)
	GetRandMeal(c *gin.Context)
}

type mealHandler struct {
	//services used by routes
	//just mealservice for now
	mealService dh.MealService
}

//creates a new server
func NewMealHandler(ms dh.MealService) MealHandler {
	return &mealHandler{
		mealService: ms,
	}
}

func NewServer() *gin.Engine {
	return gin.Default()
}
