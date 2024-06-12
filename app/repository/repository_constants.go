package repository

const (
	Email  = "email"
	TaskID = "task_id"
	UserID = "user_id"
)

// Query constants
const (
	IfTaskIdAndUserIdEqual = "task_id = ? AND user_id = ?"
)

// Repository methods
const (
	// User repository methods
	AddUserMethod 			= "AddUser"
	GetUserByEmailMethod 	= "GetUserByEmail"
	// Task repository methods
	AddTaskMethod 				= "AddTask"
	UpdateTaskMethod 			= "UpdateTask"
	GetTasksByUserIDMethod 		= "GetTasksByUserID"
	DeleteTaskMethod 			= "DeleteTask"
	IsTaskExistforUserMethod 	= "IsTaskExistforUser"
)

// Table names
const (
	Users = "users"
	Tasks = "tasks"
)