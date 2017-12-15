package http

type CreateBirdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateBirdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
