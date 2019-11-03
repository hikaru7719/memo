package entity

import (
	"github.com/jinzhu/gorm"
)

// Done represents a task user finised
type Done struct {
	gorm.Model
	Name        string
	Explanation string
	TagID       int
	UserID      string
}

// Tag units ToDo tasks or Done tasks
type Tag struct {
	gorm.Model
	Name   string
	UserID string
}
