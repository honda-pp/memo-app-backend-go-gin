package main

import (
	"github.com/honda-pp/memo-app-backend-go-gin/app/handlers"
	"github.com/honda-pp/memo-app-backend-go-gin/infrastructure/logger"

	generated "github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func main() {
	log := logger.InitLogger()

	api := handlers.UsersHandler{}

	routes := generated.NewApiHandleFunctions(&api)

	log.Printf("Server started")

	router := generated.NewRouter(routes)

	log.Fatal(router.Run(":8000"))
}
