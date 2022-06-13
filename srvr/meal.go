package srvr

import (
	"context"

	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMeal(c *gin.Context) error {
	ctx := context.Background()

	var meal *dh.Meal

	err := c.ShouldBindJSON(&meal)

	if err != nil {
		return err
	}

	h.MealService.CreateMeal(ctx, meal)

	return nil
}

//update this to show all
func (h *Handler) FindAll() ([]*dh.Meal, error) {
	ctx := context.Background()

	return h.MealService.FindAll(ctx)

}

//update this to show a random meal
func (h *Handler) RandMeal(c *gin.Context) (*dh.Meal, error) {
	//placeholder for now
	return nil, nil
}
