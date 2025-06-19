package repo

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
