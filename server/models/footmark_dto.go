package models

type AppFootmarkItem struct {
	Id       uint    `json:"id"`
	ImageUrl string  `json:"imageUrl"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
}
