package models

import (
    "time"
)

type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    Complete    bool      `json:"complete"`
}

// IsOverdue returns true if the task due date is before the current time.
func (t *Task) IsOverdue() bool {
    return t.DueDate.Before(time.Now())
}

// UpdateDueDate updates the due date of the task.
func (t *Task) UpdateDueDate(newDueDate time.Time) {
    t.DueDate = newDueDate
}

// MarkComplete marks the task as completed.
func (t *Task) MarkComplete() {
    t.Complete = true
}