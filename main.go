package main

import (
	"github.com/gorilla/mux"
	"go-mux-gorm/config"
	"go-mux-gorm/controller"
	"go-mux-gorm/model"
	"go-mux-gorm/util"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:7000")

	db := config.Open()
	db.AutoMigrate(model.Movie{})

	router := mux.NewRouter()
	router.Use(util.JwtAuthentication)
	router.HandleFunc("/movies", controller.FindAll).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.FindById).Methods("GET")
	router.HandleFunc("/movies", controller.Save).Methods("POST")
	router.HandleFunc("/movies", controller.Update).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.Delete).Methods("DELETE")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	if err := http.ListenAndServe(":7000", router); err != nil {
		panic(err.Error())
	}
}
