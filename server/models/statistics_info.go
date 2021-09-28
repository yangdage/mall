package models

type StatisticsInfo struct {
	GoodsCount   int     `json:"goodsCount"`
	OrderCount   int     `json:"orderCount"`
	Amount       float64 `json:"amount"`
	VisitorCount int     `json:"visitorCount"`
}
