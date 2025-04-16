package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"htmx-graphql-healthcare/config"
	"htmx-graphql-healthcare/handler"
	"htmx-graphql-healthcare/repository"
	"htmx-graphql-healthcare/service"
)

func main() {
	db := config.ConnectMongoDB()
	repo := repository.NewMongoRepo(db)
	svc := service.NewService(repo)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/medications", handler.GetMedicationsHandler(svc)).Methods("GET")
	r.HandleFunc("/add-medication", handler.AddMedicationHandler(svc)).Methods("POST")

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
