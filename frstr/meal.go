package frstr

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
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

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Could not create firestore client: %v", err)
	}

	defer client.Close()

	//I think this will work without explicitly setting everything?
	//otherwise see daren repo for how to set fields in a map
	_, _, err = client.Collection(mealCollection).Add(ctx, m)

	if err != nil {
		log.Fatalf("Failed to create a new meal: %v", err)
	}

	//will eventually want to return something other than nil
	//also prob don't want to fatally close the application if errors
	return nil
}

//GetAllMeals method
func (ms *mealService) GetAllMeals(ctx context.Context) ([]*dh.Meal, error) {
	//implement db logic
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Could not create firestore client: %v", err)
	}

	defer client.Close()

	var meals []*dh.Meal

	iter := client.Collection(mealCollection).Documents(ctx)

	docs, err := iter.GetAll()

	if err != nil {
		log.Fatalf("Couldn't get data from firestore: %v", err)
	}

	for _, doc := range docs {
		var m *dh.Meal

		//this should work; if not, see daren stuff
		doc.DataTo(&m)

		meals = append(meals, m)
	}

	//same deal here -- will want to return an error that propogates to the user at some point
	return meals, nil
}

//GetRandMeal method
func (ms *mealService) GetRandMeal(ctx context.Context) (*dh.Meal, error) {
	//implement db logic
	return nil, nil
}
