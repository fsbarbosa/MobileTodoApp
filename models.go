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
    Priority    string    `json:"priority"` // Added priority field
    CreatedAt   time.Time `json:"createdAt"` // Added createdAt field
}

// IsOverduereturns true if the task due date is before the current time.
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

// MarkIncomplete marks the task as not completed.
func (t *Task) MarkIncomplete() {
    t.Complete = false
}

// UpdateTaskDetails updates both the title and description of the task.
func (t *Task) UpdateTaskDetails(newTitle string, newDescription string) {
    t.Title = newTitle
    t.Description = newDescription
}

// SetPriority sets the priority level of the task.
func (t *Task) SetPriority(priority string) {
    t.Priority = priority
}