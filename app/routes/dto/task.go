package dto

import "time"

type Task struct {
	*Base
	UserID      uint   		`json:"user_id"`
	PriorityID  uint   		`json:"priority_id"`
	Title       string 		`json:"title"`
	Description string 		`json:"description"`
	Status      string 		`json:"status"`
	DueDate     time.Time 	`json:"due_date"`
}

type Tags struct {
	*Base
	TagName string `json:"tag_name"`
}

type TaskTags struct {
	*Base
	TaskID uint `json:"task_id"`
	TagID  uint `json:"tag_id"`
}

type Priority struct {
	*Base
	PriorityLevel string `json:"priority_level"`
}