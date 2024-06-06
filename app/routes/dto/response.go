package dto

type UserRegistrationResponse struct {
	UserID uint 	`json:"UserID"`
	Email  string	`json:"Email"`
	Name   string 	`json:"Name"`
}