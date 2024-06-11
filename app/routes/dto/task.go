package dto

import "time"

type PriorityLevel string

type Status string

const (
	High   PriorityLevel = "high"
	Medium PriorityLevel = "medium"
	Low    PriorityLevel = "low"
)

const (
	Completed 	Status 	= "completed"
	ToDo     	Status 	= "to_do"
	OnHold    	Status 	= "on_hold"
)

// many2many mapping with Tags
type Tasks struct {
	*Base
	TaskID      uint   		`gorm:"not null; column:task_id; primaryKey; autoIncrement;"`
	UserID      uint   		`gorm:"not null; column:user_id"` // Foreign key
	PriorityID  uint   		`gorm:"not null; column:priority_id"` // Foreign key
	Title       string 		`gorm:"not null; column:title"`
	Description string 		`gorm:"column:description"`
	Status      string 		`gorm:"not null; column:status"`
	DueDate     time.Time 	`gorm:"not null; column:due_date"`
	Tags        []Tags 		`gorm:"many2many:task_tags;"`
}

//  many2many mapping with Tasks
type Tags struct {
	*Base
	TagID    uint  		`gorm:"not null; column:tag_id; primaryKey; autoIncrement;"`
	TagName string 		`gorm:"not null; column:tag_name"`
	Tasks    []Tasks 	`gorm:"many2many:task_tags;"`
}

// join table for many2many relationship between Tasks and Tags
type TaskTags struct {
	*Base
	TaskID uint `gorm:"not null; column:task_id"`
	TagID  uint `gorm:"not null; column:tag_id"`
}

// 1:many mapping with Tasks
type Priorities struct {
	*Base
	PriorityID   	uint   	`gorm:"not null; column:priority_id; primaryKey; autoIncrement;"`
	PriorityLevel 	string 	`gorm:"not null; column:priority_level"`
	Tasks		 	[]Tasks `gorm:"foreignKey:PriorityID"`
}