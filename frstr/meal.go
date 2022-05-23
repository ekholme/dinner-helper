//an implementation that calls a firestore db
package frstr

import (
	dh "github.com/ekholme/dinner-helper"
)

type mealService struct{}

func NewMealService() dh.MealService {
	return &mealService{}
}

//TODO
//set up GCP project
//add methods to satisfy interface
