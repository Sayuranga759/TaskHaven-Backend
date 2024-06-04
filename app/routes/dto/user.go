package dto

type User struct {
	*Base
	Username string `gorm:"not null; column:username; unique"`
	Email    string `gorm:"not null; column:email; unique"`
	Password string `gorm:"not null; column:password"`
}

type JWTClaims struct {
	UserID uint 	
	Email  string 	
}