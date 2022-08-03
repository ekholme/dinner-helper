package frstr

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	dh "github.com/ekholme/dinner-helper"
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
	//this function currently throwing an error
	//see this for help https://stackoverflow.com/questions/63214521/how-to-use-firestore-documentid-name-in-query
	//i think i need to generate a new document ref

	//get a random uuid to reference against
	//ref := uuid.New().String()

	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Couldn't create firestore client: %v", err)
	}

	//ok so the solution might actually be to just add in a random id (int64 or uuid) to each document that lets me query this field

	defer client.Close()

	var meals []*dh.Meal

	//this is the 'default' method to get a random document
	// iter := client.Collection(mealCollection).Where(firestore.DocumentID, ">=", ref).OrderBy(firestore.DocumentID, firestore.Asc).Limit(1).Documents(ctx)

	iter := client.Collection(mealCollection).Where(firestore.DocumentID, ">=", "").Limit(1).Documents(ctx)

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

	//run the 'less than' query if the 'greater than' query returns 0 results
	if len(meals) == 0 {
		// iter := client.Collection(mealCollection).Where(firestore.DocumentID, "<=", ref).OrderBy(firestore.DocumentID, firestore.Asc).Limit(1).Documents(ctx)

		iter := client.Collection(mealCollection).Where(firestore.DocumentID, ">=", "").Limit(1).Documents(ctx)

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

	m = meals[0]

	//implement db logic
	return m, nil
}
