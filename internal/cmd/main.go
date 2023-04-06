package main

import (
	"github.com/gin-gonic/gin"
	user_controller "github.com/mauFade/fit-force/internal/infra/controller/user"
	"github.com/mauFade/fit-force/internal/loaders"
)

func init() {
	loaders.GetEnvironmentVariables()
	loaders.ConnectToDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/users", user_controller.CreateUserController)

	router.Run()
}
