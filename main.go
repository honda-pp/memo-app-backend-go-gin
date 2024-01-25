package main

import (
	"github.com/honda-pp/memo-app-backend-go-gin/app/handlers"
	"github.com/honda-pp/memo-app-backend-go-gin/infrastructure/logger"

	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func main() {
	log := logger.InitLogger()

	memoApi := handlers.NewMemoHandler()

	userApi := handlers.NewUsersHandler()

	routes := generated.NewApiHandleFunctions(memoApi, userApi)

	log.Printf("Server started")

	router := generated.NewRouter(routes)

	log.Fatal(router.Run(":8000"))
}
