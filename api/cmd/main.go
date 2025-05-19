package main

import (
	"gin-auth-boilerplate/config"
	"gin-auth-boilerplate/internal/routes"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	if err := config.InitDB(); err != nil {
		logger.Error("error initializing database: %v", err)
		return
	}

	router := routes.InitializeRoutes()
	router.Run()
}
