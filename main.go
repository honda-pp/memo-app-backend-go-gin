package main

import (
	"github.com/honda-pp/memo-app-backend-go-gin/app/handlers"
	"github.com/honda-pp/memo-app-backend-go-gin/app/repositories"
	"github.com/honda-pp/memo-app-backend-go-gin/app/usecases"
	"github.com/honda-pp/memo-app-backend-go-gin/infrastructure/database"
	"github.com/honda-pp/memo-app-backend-go-gin/infrastructure/logger"

	"github.com/honda-pp/memo-app-backend-go-gin/generated"
)

func main() {
	log := logger.InitLogger()
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userApi := handlers.NewUsersHandler(userUsecase)

	memoRepository := repositories.NewMemoRepository(db)
	memoUsecase := usecases.NewMemoUsecase(memoRepository)
	memoApi := handlers.NewMemoHandler(memoUsecase)

	routes := generated.NewApiHandleFunctions(memoApi, userApi)

	log.Printf("Server started")

	router := generated.NewRouter(routes)

	log.Fatal(router.Run(":8000"))
}
