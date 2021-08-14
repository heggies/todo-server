package presenter

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
}
