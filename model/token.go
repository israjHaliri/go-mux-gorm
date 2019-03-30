package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserId uint
	jwt.StandardClaims
}
