package frstr

import (
	"context"
	"log"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
	dh "github.com/ekholme/dinner-helper"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
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

	_, _, err = client.Collection(mealCollection).Add(ctx, m)

	if err != nil {
		log.Fatalf("Failed to create a new meal: %v", err)
	}

	return nil
}

//GetAllMeals method
func (ms *mealService) GetAllMeals(ctx context.Context) ([]*dh.Meal, error) {

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
		return nil, err
	}

	for _, doc := range docs {
		var m *dh.Meal

		doc.DataTo(&m)

		meals = append(meals, m)
	}

	return meals, nil
}

//GetRandMeal method
func (ms *mealService) GetRandMeal(ctx context.Context) (*dh.Meal, error) {
	//set the random seed
	rand.Seed(time.Now().UnixNano())

	//get a random uuid to reference against
	ref := uuid.New().String()

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Couldn't create firestore client: %v", err)
	}

	defer client.Close()

	var meals []*dh.Meal

	iter := client.Collection(mealCollection).Where("ID", ">=", ref).OrderBy("ID", 1).Limit(1).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var m *dh.Meal

		doc.DataTo(&m)

		meals = append(meals, m)
	}

	//wrap around from the other side if the previous iteration returns nothing
	if len(meals) == 0 {

		iter := client.Collection(mealCollection).Where("ID", "<=", ref).OrderBy("ID", 1).Limit(1).Documents(ctx)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}

			var m *dh.Meal

			doc.DataTo(&m)

			meals = append(meals, m)
		}
	}

	var m *dh.Meal

	//just in case this returns multiple, only take the first meal
	m = meals[0]

	return m, nil
}
