package gift

import (
	"time"
)

type Gift struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	ProductID  int64     `json:"product_id"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
