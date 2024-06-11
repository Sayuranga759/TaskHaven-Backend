package dto

import "github.com/golang-jwt/jwt/v4"

type Users struct {
	*Base
	UserID   	uint   	`gorm:"not null; column:user_id; primaryKey; autoIncrement;"`
	Name     	string 	`gorm:"not null; column:name"`
	Email    	string 	`gorm:"not null; column:email; unique"`
	Password 	string 	`gorm:"not null; column:password"`
	Tasks		[]Tasks `gorm:"foreignKey:UserID"`
}

type JWTClaims struct {
	UserID 	uint
	Email  	string
	Name  	string
	jwt.RegisteredClaims
}