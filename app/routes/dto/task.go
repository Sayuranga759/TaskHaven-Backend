package dto

import "time"

type Task struct {
	*Base
	UserID      uint   		`gorm:"not null; column:user_id"` // Foreign key
	PriorityID  uint   		`gorm:"not null; column:priority_id"` // Foreign key
	Title       string 		`gorm:"not null; column:title"`
	Description string 		`gorm:"column:description"`
	Status      string 		`gorm:"not null; column:status"`
	DueDate     time.Time 	`gorm:"not null; column:due_date"`
}

type Tags struct {
	*Base
	TagName string `gorm:"not null; column:tag_name"`
}

type TaskTags struct {
	*Base
	TaskID uint `gorm:"not null; column:task_id"`
	TagID  uint `gorm:"not null; column:tag_id"`
}

type Priority struct {
	*Base
	PriorityLevel string `gorm:"not null; column:priority_level"`
}