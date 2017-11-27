package models

type Product struct {
	ID      int
	Size    string
	Color   string
	Price   float64
	ImageID string
	InStock int
	GroupID int
}
