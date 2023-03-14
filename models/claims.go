package models

import "github.com/golang-jwt/jwt"

type AppClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
