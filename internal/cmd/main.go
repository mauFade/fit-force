package main

import "github.com/mauFade/fit-force/internal/loaders"

func init() {
	loaders.GetEnvironmentVariables()
	loaders.ConnectToDatabase()
}
