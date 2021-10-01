package models

type OrderList struct {
	Id            uint    `json:"id"`
	PaymentStatus int     `json:"paymentStatus"`
	Status        string  `json:"status"`
	UserName      string  `json:"userName"`
	TotalPrice    float64 `json:"totalPrice"`
	Created       string  `json:"created"`
}
