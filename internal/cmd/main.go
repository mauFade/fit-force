package main

import (
	"github.com/gin-gonic/gin"
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

	// Workout routes
	router.GET("/workouts", middleware.AuthMiddleware(), workout_controller.ListWorkoutsController)

	router.Run()
}
