package models

type WebStatisticsInfo struct {
	GoodsCount   int     `json:"goodsCount"`
	OrderCount   int     `json:"orderCount"`
	Amount       float64 `json:"amount"`
}

type WebTodayOrderInfo struct {
	Data [6]int `json:"data"`
}

type WebWeekOrderInfo struct {
	Orders [7]int     `json:"orders"`
	Amount [7]float64 `json:"amount"`
}