package dto

import "time"

type UserRegistrationRequest struct {
	Name     string `json:"Name" validate:"required,max=50,alphaWithSpace"`
	Email    string `json:"Email" validate:"required,max=100,email"`
	Password string `json:"Password" validate:"required,min=8,password"`
}

type LoginRequest struct {
	Email    string `json:"Email" validate:"required,max=100,email"`
	Password string `json:"Password" validate:"required,min=8"`
}

type ValidateTokenRequest struct {
	AuthString string `json:"AuthString"`
}

type CreateTaskRequest struct {
	UserID      uint   		`json:"UserID" `
	PriorityID  uint   		`json:"PriorityID"`
	Title       string 		`json:"Title" validate:"required,max=100"`
	Description string 		`json:"Description" validate:"max=250"`
	Status      Status 		`json:"Status" validate:"required,oneof=completed to_do on_hold"`
	DueDate     time.Time 	`json:"DueDate" validate:"required,timestamp"`
}

type UpdateTaskRequest struct {
	TaskID      uint   		`json:"TaskID" validate:"required"`
	*CreateTaskRequest
}

type DeleteTaskRequest struct {
	TaskID uint `json:"TaskID" validate:"required"`
	UserID uint `json:"UserID"`
}