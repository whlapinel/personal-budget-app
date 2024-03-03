package main

import (
	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"username"`
	jwt.RegisteredClaims
}


