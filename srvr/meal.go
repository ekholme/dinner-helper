package srvr

import (
	"context"
	"net/http"

	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

//an interface to take care of meal stuff
type MealHandler interface {
	CreateMeal(c *gin.Context)
	GetAllMeals(c *gin.Context)
	GetRandMeal(c *gin.Context)
}

type mealHandler struct {
	//services used by routes
	//just mealservice for now
	mealService dh.MealService
}

//create a new instance of mealHandler
func NewMealHandler(ms dh.MealService) MealHandler {
	return &mealHandler{
		mealService: ms,
	}
}

//implement methods for mealHandler
func (mh mealHandler) CreateMeal(c *gin.Context) {
	ctx := context.Background()

	var meal *dh.Meal

	err := c.ShouldBindJSON(&meal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = mh.mealService.CreateMeal(ctx, meal)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "meal saved!"})
}

func (mh mealHandler) GetAllMeals(c *gin.Context) {
	ctx := context.Background()

	meals, err := mh.mealService.GetAllMeals(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dares not retrieved"})
		return
	}

	data := gin.H{
		"meals": meals,
	}

	c.JSON(http.StatusOK, data)
}

func (mh mealHandler) GetRandMeal(c *gin.Context) {
	//todo
}
