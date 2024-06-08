package dto

import "github.com/golang-jwt/jwt/v4"

type User struct {
	*Base
	Name     string `gorm:"not null; column:name"`
	Email    string `gorm:"not null; column:email; unique"`
	Password string `gorm:"not null; column:password"`
}

type JWTClaims struct {
	UserID 	uint
	Email  	string
	Name  	string
	jwt.RegisteredClaims
}