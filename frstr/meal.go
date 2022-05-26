//an implementation that calls a firestore db
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

type mealService struct{}

func NewMealService() dh.MealService {
	return &mealService{}
}

func newContext() context.Context {
	return context.Background()
}

func (*mealService) CreateMeal(ctx context.Context, m *dh.Meal) error {
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Couldn't create Firestore client: %v", err)
		return err
	}

	defer client.Close()

	_, _, err = client.Collection(mealCollection).Add(ctx, map[string]interface{}{
		"Name":       m.Name,
		"Time":       m.Time,
		"Notes":      m.Notes,
		"Difficulty": m.Difficulty,
		"Link":       m.Link,
		"Prep":       m.Prep,
		"Protein":    m.Protein,
	})

	if err != nil {
		log.Fatalf("Failed to create a new meal: %v", err)
		return err
	}

	return nil

}

func (*mealService) FindAll(ctx context.Context) ([]*dh.Meal, error) {
	//add this
}

func (*mealService) RandMeal(ctx context.Context) (*dh.Meal, error) {
	//add this
}

//TODO
//set up GCP project
//add methods to satisfy interface
