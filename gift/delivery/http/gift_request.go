package http

type CreateGiftRequest struct {
	ToUserID  int64  `json:"to_user_id"`
	ProductID int64  `json:"product_id"`
	Message   string `json:"message"`
}

type UpdateGiftRequest struct {
	ProductID int64  `json:"product_id"`
	Message   string `json:"message"`
}
