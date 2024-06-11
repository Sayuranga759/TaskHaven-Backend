package dto

import "time"

type UserRegistrationResponse struct {
	UserID uint   `json:"UserID"`
	Email  string `json:"Email"`
	Name   string `json:"Name"`
}

type LoginResponse struct {
	UserID      uint   `json:"UserID"`
	AccessToken string `json:"AccessToken"`
}

type CreateTaskResponse struct {
	PriorityID  uint      `json:"PriorityID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Status      Status    `json:"Status"`
	DueDate     time.Time `json:"DueDate"`
}