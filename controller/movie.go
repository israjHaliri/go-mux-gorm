package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-mux-gorm/model"
	"go-mux-gorm/service"
	"go-mux-gorm/util"
	"net/http"
	"strconv"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	movies := service.FindAll()

	if movies == nil {
		util.ResponseWithError(w, http.StatusInternalServerError, "Not Found")
		return
	}

	util.ResponseWithJson(w, http.StatusOK, movies)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	movie := service.FindById(id)

	if movie == nil {
		util.ResponseWithError(w, http.StatusBadRequest, "Invalid Movie ID")

		return
	}

	util.ResponseWithJson(w, http.StatusOK, movie)
}

func Save(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie model.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")

		return
	}

	if _, err := service.Save(movie); err != nil {
		util.ResponseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	util.ResponseWithJson(w, http.StatusCreated, movie)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var movie model.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")

		return
	}

	if _, err := service.Update(movie); err != nil {
		util.ResponseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	util.ResponseWithJson(w, http.StatusCreated, movie)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	if err := service.Delete(id); err != nil {
		util.ResponseWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	util.ResponseWithJson(w, http.StatusCreated, "Deleted")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var account model.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		util.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")

		return
	}

	account, status := service.Login(account)

	if status {
		util.ResponseWithJson(w, http.StatusOK, account)
	} else {
		util.ResponseWithJson(w, http.StatusNotFound, "Username or passowrd is wrong")
	}
}
