package models

import "time"

type Todo struct {
	ID          uint      `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Done        bool      `json:"done,omitempty"`
	CreatedAt   time.Time `json:"created_at" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" json:"updated_at"`
}
