package models

// 商品信息传输模型
type WebProductInfo struct {
	Id           uint    `json:"id"`
	CategoryId   uint    `json:"categoryId"`
	Kind         int     `json:"kind"`
	Title        string  `json:"title"`
	BrandId      uint    `json:"brandId"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Amount       int     `json:"amount"`
	Sales        int     `json:"sales"`
	ImageUrl     string  `json:"imageUrl"`
	SendAddress  string  `json:"sendAddress"`
	ParcelType   string  `json:"parcelType"`
	SalesService string  `json:"salesService"`
}

// 商品项传输模型
type WebProductItem struct {
	Id       uint    `gorm:"primaryKey" json:"id"`
	ImageUrl string  `gorm:"imageUrl"   json:"imageUrl"`
	Title    string  `gorm:"title"      json:"title"`
	Name     string  `gorm:"name"       json:"name"`
	Price    float64 `gorm:"price"      json:"price"`
}

// 商品项传输模型
type AppProductItem struct {
	Id       uint    `gorm:"primaryKey" json:"id"`
	ImageUrl string  `gorm:"imageUrl"   json:"imageUrl"`
	Title    string  `gorm:"title"      json:"title"`
	Price    float64 `gorm:"price"      json:"price"`
}

// 商品列表传输模型
type WebProductList struct {
	Id       uint    `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Amount   int     `json:"amount"`
	ImageUrl string  `json:"imageUrl"`
	Sales    int     `json:"sales"`
	Status   int     `json:"status"`
	Created  string  `json:"created"`
}

// 微信小程序，商品列表传输模型
type AppProductList struct {
	Id       uint    `json:"id"`
	ImageUrl string  `json:"imageUrl"`
	Title    string  `json:"title"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

// 微信小程序，商品信息传输模型
type AppProductDetail struct {
	Id           uint    `json:"id"`
	Title        string  `json:"title"`
	Price        float64 `json:"price"`
	Amount       int     `json:"amount"`
	Sales        int     `json:"sales"`
	ImageUrl     string  `json:"imageUrl"`
	SendAddress  string  `json:"sendAddress"`
	ParcelType   string  `json:"parcelType"`
	SalesService string  `json:"salesService"`
}
