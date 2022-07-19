package frstr

import (
	"context"

	dh "github.com/ekholme/dinner-helper"
)

const (
	projectID      = "fitz-dinner-helper"
	mealCollection = "meals"
)

//create an empty struct to implement MealService interface
type mealService struct{}

//constructor function
func NewMealService() dh.MealService {
	return &mealService{}
}

//CreateMeal method
func (ms *mealService) CreateMeal(ctx context.Context, m *dh.Meal) error {
	//implement db logic
	return nil
}

//GetAllMeals method
func (ms *mealService) GetAllMeals(ctx context.Context) ([]*dh.Meal, error) {
	//implement db logic
	return nil, nil
}

//GetRandMeal method
func (ms *mealService) GetRandMeal(ctx context.Context) (*dh.Meal, error) {
	//implement db logic
	return nil, nil
}
