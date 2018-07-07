package http

type CreateBirdRequest struct {
	Name          string `json:"name"`	
	Color       	string `json:"color"`
	Description 	string `json:"description"`
}

type UpdateBirdRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
