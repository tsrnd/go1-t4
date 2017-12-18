package class

import "time"

type Class struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	StudentNumber int       `json:"student_number"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
