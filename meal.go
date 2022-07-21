package dh

import (
	"context"
)

//Meal represents a meal to include in dinner helper
type Meal struct {
	Name       string `json:"name" binding:"required"`
	Time       int64  `json:"time"`
	Notes      string `json:"notes"`
	Difficulty int64  `json:"difficulty" binding:"gte=1,lte=5"`
	Link       string `json:"link"`    //see if there's a url validator
	Protein    string `json:"protein"` //represent the main protein in the meal
}

type MealService interface {
	CreateMeal(ctx context.Context, m *Meal) error     //save a new meal
	GetAllMeals(ctx context.Context) ([]*Meal, error) //get all of the meals
	GetRandMeal(ctx context.Context) (*Meal, error)    //get a random meal
	//add UpdateMeal here at some point
}
