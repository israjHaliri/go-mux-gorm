package service

import (
	"go-mux-gorm/model"
	"testing"
)

func TestLogin(t *testing.T) {
	var account model.Account
	account.Email = "israj.haliri@gmail.com"
	account.Password = "12345678"

	_, status := Login(account)

	if !status {
		t.Fail()
	}
}
