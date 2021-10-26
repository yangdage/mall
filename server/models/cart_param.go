package models

type AppCartParam struct {
	ProductId uint   `form:"productId" json:"productId"`
	UserId    string `form:"userId"    json:"userId"`
}
