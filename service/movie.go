package service

import (
	"go-mux-gorm/config"
	"go-mux-gorm/model"
)

func FindAll() []model.Movie {
	db := config.Open()
	defer db.Close()

	movie := []model.Movie{}

	db.Find(&movie)

	return movie
}

func FindById(id int) interface{} {
	db := config.Open()
	defer db.Close()

	movie := model.Movie{}

	err := db.Where("id = ?", id).First(&movie).Error

	if err != nil {
		return nil
	}

	return movie
}

func Save(movie model.Movie) (model.Movie, error) {
	db := config.Open()
	defer db.Close()

	err := db.Create(&movie).Error

	return movie, err
}

func Update(movie model.Movie) (model.Movie, error) {
	db := config.Open()
	defer db.Close()

	err := db.Save(&movie).Error

	return movie, err
}

func Delete(id int) (error) {
	db := config.Open()
	defer db.Close()

	err := db.Delete(&model.Movie{}, "id = ?", id).Error

	return err
}
