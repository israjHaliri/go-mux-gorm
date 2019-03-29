package model

type Movie struct {
	ID    uint `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Price int `json:"price"`
}
