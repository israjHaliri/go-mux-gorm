package main

import (
	"github.com/gorilla/mux"
	"go-mux-gorm/config"
	"go-mux-gorm/controller"
	"go-mux-gorm/model"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:7000")

	db := config.Open()
	db.AutoMigrate(model.Movie{})

	r := mux.NewRouter()
	r.HandleFunc("/movies", controller.FindAll).Methods("GET")
	r.HandleFunc("/movies/{id}",controller. FindById).Methods("GET")
	r.HandleFunc("/movies", controller.Save).Methods("POST")
	r.HandleFunc("/movies", controller.Update).Methods("PUT")
	r.HandleFunc("/movies/{id}", controller.Delete).Methods("DELETE")

	if err := http.ListenAndServe(":7000", r); err != nil {
		panic(err.Error())
	}
}
