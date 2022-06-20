package srvr

import (
	"context"
	"net/http"

	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

//I need to think about the return value for these
//they shouldn't return an error; instead, they should return nothing
//and instead create html/json that's appropriate for a given error/success
//may also want to define a handler interface
func (mh *mealHandler) CreateMeal(c *gin.Context) {
	ctx := context.Background()

	var meal *dh.Meal

	err := c.ShouldBindJSON(&meal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = mh.mealService.CreateMeal(ctx, meal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"status": "meal saved"})

}

//update this to show all
func (mh *mealHandler) FindAllMeals(c *gin.Context) {
	ctx := context.Background()

	meals, err := mh.mealService.FindAllMeals(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dares not retrieved"})
		return
	}

	data := gin.H{
		"meals": meals,
	}

	c.JSON(http.StatusOK, data)

}

//update this to show a random meal
func (mh *mealHandler) GetRandMeal(c *gin.Context) {
	//placeholder for now
}
