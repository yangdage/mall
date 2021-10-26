package models

type AppFootmarkParam struct {
	ProductId uint   `form:"productId" json:"productId"`
	UserId    string `form:"userId"    json:"userId"`
}

