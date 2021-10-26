package models

// 商品表单参数模型
type WebProductFormParam struct {
	CategoryId   uint    `json:"categoryId"   binding:"required,gt=0"`
	Kind         int     `json:"kind"         binding:"required,oneof=1 2"`
	Title        string  `json:"title"        binding:"required"`
	BrandId      uint    `json:"brandId"      binding:"required,gt=0"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"        binding:"required,gt=0"`
	Amount       int     `json:"amount"       binding:"required,gt=0"`
	ImageUrl     string  `json:"imageUrl"     binding:"required"`
	SendAddress  string  `json:"sendAddress"  binding:"required"`
	ParcelType   string  `json:"parcelType"   binding:"required"`
	SalesService string  `json:"salesService" binding:"required"`
	CreatorId    uint    `json:"creatorId"    binding:"required,gt=0"`
	Status       int     `json:"status"       binding:"required,oneof=1 2"`
}

// 商品更新参数模型
type WebProductUpdateParam struct {
	Id           uint    `json:"id"           binding:"omitempty,gt=0"`
	CategoryId   uint    `json:"categoryId"   binding:"omitempty,gt=0"`
	Kind         int     `json:"kind"         binding:"omitempty,oneof=1 2"`
	Title        string  `json:"title"        binding:"omitempty"`
	BrandId      uint    `json:"brandId"      binding:"omitempty,gt=0"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"        binding:"omitempty,gt=0"`
	Amount       int     `json:"amount"       binding:"omitempty,gt=0"`
	ImageUrl     string  `json:"imageUrl"     binding:"omitempty"`
	SendAddress  string  `json:"sendAddress"  binding:"omitempty"`
	ParcelType   string  `json:"parcelType"   binding:"omitempty"`
	SalesService string  `json:"salesService" binding:"omitempty"`
	CreatorId    uint    `json:"creatorId"    binding:"omitempty,gt=0"`
	Status       int     `json:"status"       binding:"omitempty"`
}

// 商品查询参数模型
type WebProductQueryParam struct {
	Page       Page
	Id         uint   `form:"id"           binding:"omitempty,gt=0"`
	CategoryId uint   `form:"categoryId"   binding:"omitempty,gt=0"`
	Kind       int    `form:"kind"         binding:"omitempty,oneof=1 2"`
	Title      string `form:"title"        binding:"omitempty"`
	BrandId    uint   `form:"brandId"      binding:"omitempty"`
	CreatorId  uint   `form:"creatorId"    binding:"omitempty,gt=0"`
	Status     int    `form:"status"       binding:"omitempty,oneof=1 2"`
}
