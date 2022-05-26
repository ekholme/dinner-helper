package dh

import "context"

//Meal represents a meal to include in dinner helper
type Meal struct {
	Name       string `json:"name" binding:"required"`
	Time       int    `json:"time"`
	Notes      string `json:"notes"`
	Difficulty int    `json:"difficulty" binding:"gte=1,lte=3"`
	Link       string `json:"link"`    //see if there's a url validator
	Prep       string `json:"prep"`    //this to represent stovetop, crockpot, oven, etc
	Protein    string `json:"protein"` //represent the main protein in the meal
}

type MealService interface {
	CreateMeal(ctx context.Context, m *Meal) error //save a new meal
	FindAll() ([]Meal, error)                      //get all of the meals
	RandMeal() (*Meal, error)                      //get a random meal
	//add UpdateMeal here at some point
}
