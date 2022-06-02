package gn

import (
	"context"

	dh "github.com/ekholme/dinner-helper"
	"github.com/ekholme/dinner-helper/frstr"
	"github.com/gin-gonic/gin"
)

type MealController interface {
	CreateMeal(c *gin.Context) error
	FindAll() ([]*dh.Meal, error)
}

type controller struct {
	meal frstr.MealService
}

func New(m frstr.MealService) MealController {
	return &controller{
		meal: m,
	}
}

func (cn *controller) FindAll() ([]*dh.Meal, error) {
	ctx := context.Background()

	return cn.meal.FindAll(ctx)
}

func (cn *controller) CreateMeal(c *gin.Context) error {
	ctx := context.Background()

	var meal *dh.Meal

	err := c.ShouldBindJSON(&meal)

	if err != nil {
		return err
	}

	cn.meal.CreateMeal(ctx, meal)

	return nil
}
