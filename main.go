package main

import (
	"log"

	sw "github.com/honda-pp/memo-app-backend-go-gin/generated/go"
)

func main() {
	routes := sw.ApiHandleFunctions{}

	log.Printf("Server started")

	router := sw.NewRouter(routes)

	log.Fatal(router.Run(":8080"))
}
