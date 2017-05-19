package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claims : new user detail goes here
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
