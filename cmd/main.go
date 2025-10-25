package main

import (
	"log"
	"net/http"

	"server/cmd/handlers"
	"server/cmd/repositories"
	"server/cmd/routes"
)

func main() {
	//bookRepo := repositories.NewInMemoryBookRepository()
	bookRepo, err := repositories.NewSqliteBookRepository("books.db")
	if err != nil {
		log.Fatal(err)
	}
	bookHandler := handlers.NewBookHandler(bookRepo)
	routes.SetupPingRoute()
	routes.SetupBookRoutes(bookHandler)
	log.Println("Server is running on port 3030")
	http.ListenAndServe(":3030", nil)
}
