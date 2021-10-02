package models

type WeekOrderInfo struct {
	Orders [7]int     `json:"orders"`
	Amount [7]float64 `json:"amount"`
}
