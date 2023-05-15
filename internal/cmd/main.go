package main

import (
	"github.com/gin-gonic/gin"
	trainingplan_controller "github.com/mauFade/fit-force/internal/infra/controller/training-plan"
	user_controller "github.com/mauFade/fit-force/internal/infra/controller/user"
	workout_controller "github.com/mauFade/fit-force/internal/infra/controller/workout"
	"github.com/mauFade/fit-force/internal/infra/middleware"
	"github.com/mauFade/fit-force/internal/loaders"
)

func init() {
	loaders.GetEnvironmentVariables()
	loaders.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	// User routes
	router.POST("/users", user_controller.CreateUserController)

	router.POST("/login", user_controller.AuthenticateController)

	// Workout routes
	router.POST("/workouts", middleware.AuthMiddleware(), workout_controller.CreateWorkoutController)
	router.GET("/workouts", middleware.AuthMiddleware(), workout_controller.ListWorkoutsController)

	router.POST("/training-plan", middleware.AuthMiddleware(), trainingplan_controller.CreateTrainingPlanController)
	router.GET("/training-plan", middleware.AuthMiddleware(), trainingplan_controller.ListTrainingPlanController)

	router.Run()
}
