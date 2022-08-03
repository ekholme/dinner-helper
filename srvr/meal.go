package srvr

import (
	"context"
	"net/http"

	dh "github.com/ekholme/dinner-helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//an interface to take care of meal stuff
type MealHandler interface {
	CreateMeal(c *gin.Context)
	GetAllMeals(c *gin.Context)
	GetRandMeal(c *gin.Context)
}

type mealHandler struct {
	//services used by routes
	//just mealservice for now
	mealService dh.MealService
}

//create a new instance of mealHandler
func NewMealHandler(ms dh.MealService) MealHandler {
	return &mealHandler{
		mealService: ms,
	}
}

//register meal routes
func (s *Server) registerMealRoutes() {

	//path to get all meals
	s.router.GET("/meal", s.mh.GetAllMeals)

	//path to create a new meal
	s.router.POST("/meal", s.mh.CreateMeal)

	//getting a random meal
	s.router.GET("/rand_meal", s.mh.GetRandMeal)
}

//implement methods for mealHandler
func (mh mealHandler) CreateMeal(c *gin.Context) {
	ctx := context.Background()

	var meal *dh.Meal

	err := c.ShouldBindJSON(&meal)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()

	meal.ID = id

	err = mh.mealService.CreateMeal(ctx, meal)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "meal saved!"})
}

func (mh mealHandler) GetAllMeals(c *gin.Context) {
	ctx := context.Background()

	meals, err := mh.mealService.GetAllMeals(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "dares not retrieved"})
		return
	}

	data := gin.H{
		"meals": meals,
	}

	c.HTML(http.StatusOK, "all_meals", data)
}

func (mh mealHandler) GetRandMeal(c *gin.Context) {

	ctx := context.Background()

	m, err := mh.mealService.GetRandMeal(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})

		return
	}

	//RESUME HERE

	data := gin.H{
		"meal": m,
	}

	c.JSON(http.StatusOK, data)

	//c.JSON(http.StatusOK, gin.H{"msg": "success"})

	//c.HTML(http.StatusOK, "oops", gin.H{})
}
