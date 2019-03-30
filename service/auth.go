package service

import (
	"github.com/dgrijalva/jwt-go"
	"go-mux-gorm/model"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func Login(account model.Account) (model.Account, bool) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(account.Password))

	if account.Email != "israj.haliri@gmail.com" || err != nil {
		return account, false
	}

	//Create new JWT token for the newly registered account
	tk := &model.Token{UserId: 12345678}

	//one minute
	ttl := 60 * time.Second
	tk.ExpiresAt = time.Now().UTC().Add(ttl).Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	account.Password = ""
	account.Token = tokenString

	return account, true
}
