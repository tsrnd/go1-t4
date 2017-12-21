package http

type CreateBookRequest struct {
	Name       string `json:"name"`
	Content string `json:"content"`
}

type UpdateBookRequest struct {
	Name       string `json:"name"`
	Content string `json:"content"`
}
