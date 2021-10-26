package models

type AppCartItem struct {
	Id       uint    `gorm:"primaryKey" json:"id"`
	ImageUrl string  `gorm:"image_url"  json:"imageUrl"`
	Title    string  `gorm:"title"      json:"title"`
	Price    float64 `gorm:"price"      json:"price"`
}

type AppCartInfo struct {
	CartItem   []AppCartItem `json:"cartItem"`
	TotalPrice float64       `json:"totalPrice"`
}
