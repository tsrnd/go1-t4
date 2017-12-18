package http

type CreateBirdRequest struct {
	Name          string `json:"name"`	
	Color       	string `json:"color"`
	Description 	string `json:"description"`
}
