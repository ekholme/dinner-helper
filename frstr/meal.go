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

func (ms *mealService) CreateMeal(ctx context.Context, m *dh.Meal) error {
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
		"Protein":    m.Protein,
	})

	if err != nil {
		log.Fatalf("Failed to create a new meal: %v", err)
		return err
	}

	return nil

}

func (ms *mealService) FindAllMeals(ctx context.Context) ([]*dh.Meal, error) {
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Couldn't create firestore client: %v", err)
	}

	defer client.Close()

	var meals []*dh.Meal

	iter := client.Collection(mealCollection).Documents(ctx)

	docs, err := iter.GetAll()

	if err != nil {
		log.Fatalf("Couldn't get meals from firestore: %v", err)
	}

	for _, doc := range docs {
		m := &dh.Meal{
			Name:       doc.Data()["Name"].(string),
			Time:       doc.Data()["Time"].(int64),
			Notes:      doc.Data()["Notes"].(string),
			Difficulty: doc.Data()["Difficulty"].(int64),
			Link:       doc.Data()["Link"].(string),
			Protein:    doc.Data()["Protein"].(string),
		}

		meals = append(meals, m)
	}

	return meals, nil
}

func (*mealService) GetRandMeal(ctx context.Context) (*dh.Meal, error) {
	//placeholder for now
	return nil, nil
}

//TODO
//write randmeal method
