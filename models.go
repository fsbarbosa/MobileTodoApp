package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}

// IsOverdue returns true if the task due date is before the current time.
func (t *Task) IsOverdue() bool {
	return time.Now().After(t.DueDate)
}