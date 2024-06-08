package dto

type UserRegistrationRequest struct {
	Name		string	`json:"Name" validate:"required,max=50,alpha"`
	Email		string	`json:"Email" validate:"required,max=100,email"`
	Password	string	`json:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Email		string	`json:"Email" validate:"required,max=100,email"`
	Password	string	`json:"Password" validate:"required,min=8"`
}