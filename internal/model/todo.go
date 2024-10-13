package model

import "time"

// Todo is the model for the todo endpoint.
type Todo struct {
	ID           int `gorm:"primaryKey"`
	Task         string
	Description  string
	Status       Status
	Priority     int
	PriorityTask string    `gorm:"default:'medium'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

// PriorityMap is the todo priority map
var PriorityMap = map[string]int{
	"to-do":       1,
	"in-progress": 2,
	"completed":   3,
}

// CreateRequest is the request parameter for creating a new todo
type CreateRequest struct {
	TaskTitle    string `json:"task_title" validate:"required"`
	Description  string `json:"description"`
	PriorityTask string `json:"priority_task" validate:"required"`
	Status       Status `json:"status" validate:"required"`
}

// NewTodo returns a new instance of the todo model.
func NewTodo(task *CreateRequest) *Todo {
	return &Todo{
		Task:         task.TaskTitle,
		Description:  task.Description,
		PriorityTask: task.PriorityTask,
		Status:       task.Status,
	}
}

// NewUpdateTodo returns a new instance of the todo model for updating.
func NewUpdateTodo(id int, task string, status Status) *Todo {
	return &Todo{
		ID:     id,
		Task:   task,
		Status: status,
	}
}

// UpdateRequest is the request parameter for updating a todo
type UpdateRequest struct {
	UpdateRequestBody
	UpdateRequestPath
}

// UpdateRequestBody is the request body for updating a todo
type UpdateRequestBody struct {
	TaskTitle    string `json:"task_title,omitempty"`
	Description  string `json:"description,omitempty"`
	Status       Status `json:"status,omitempty"`
	PriorityTask string `json:"priority_task,omitempty"`
}

// UpdateRequestPath is the request parameter for updating a todo
type UpdateRequestPath struct {
	ID int `param:"id" validate:"required"`
}

// Status is the status of the task.
type Status string

const (
	// Created is the status for a created task.
	Created = Status("created")
	// Processing is the status for a processing task.
	Processing = Status("processing")
	// Done is the status for a done task.
	Done = Status("done")
)

// StatusMap is a map of task status.
var StatusMap = map[Status]bool{
	Created:    true,
	Processing: true,
	Done:       true,
}
