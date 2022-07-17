package dh

type Meal struct {
	Name       string `json:"name" binding:"required"`
	Time       int64  `json:"time"`
	Notes      string `json:"notes"`
	Difficulty int64  `json:"difficulty" binding:"gte=1,lte=5"`
	Link       string `json:"link"`
	Protein    string `json:"protein"`
}

type MealService interface {
	CreateMeal(ctx context.Context, m *Meal) error
	GetAllMeals(ctx context.Context) ([]*Meal, error)
	GetRandMeal(ctx context.Context) (*Meal, error)
	//add UpdateMeal at some point
}
