package book

import "time"

// Book struct
type Book struct {
	ID          int64     `json:"id"`
	Name       string    `json:"name"`
	Content string    `json:"content"`
	UserID      int64     `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
