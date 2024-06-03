package dto

type User struct {
	*Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserID uint 	`json:"user_id"`
	Email  string 	`json:"email"`
}