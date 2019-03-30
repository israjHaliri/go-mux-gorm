package model

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
