package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mux-gorm/model"
	"go-mux-gorm/service"
	"net/http"
	"strconv"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	movies := service.FindAll()

	if movies == nil {
		responseWithError(w, http.StatusInternalServerError, "Not Found")
		return
	}

	responseWithJson(w, http.StatusOK, movies)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	movie := service.FindById(id)

	if movie == nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Movie ID")

		return
	}

	responseWithJson(w, http.StatusOK, movie)
}

func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie model.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")

		return
	}

	if _, err := service.Save(movie); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	responseWithJson(w, http.StatusCreated, movie)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie model.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")

		return
	}

	if _, err := service.Update(movie); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	responseWithJson(w, http.StatusCreated, movie)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	if err := service.Delete(id); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	responseWithJson(w, http.StatusCreated, "Deleted")
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}
