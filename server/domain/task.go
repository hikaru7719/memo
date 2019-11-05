package domain

import (
	"github.com/jinzhu/gorm"
)

// Status represetns task's status
type Status string

const (
	// Unknown is status not applicative bellow status.
	Unknown Status = "Unknown"
	// Todo is status for tasks of first steps
	Todo Status = "ToDo"
	// Progress is status for tasks in progress
	Progress Status = "Progress"
	// Done is status for finished tasks
	Done Status = "Done"
)

// Task represents a task user finised
type Task struct {
	gorm.Model
	Name        string
	Explanation string
	TagID       uint
	Status      Status
	UserID      string
}

// Tag units ToDo tasks or Done tasks
type Tag struct {
	gorm.Model
	Name   string
	UserID string
}
