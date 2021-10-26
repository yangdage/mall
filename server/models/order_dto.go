package models

// 订单列表传输模型
type WebOrderList struct {
	Id            uint    `json:"id"`
	PaymentStatus int     `json:"paymentStatus"`
	Status        string  `json:"status"`
	UserName      string  `json:"userName"`
	TotalPrice    float64 `json:"totalPrice"`
	Created       string  `json:"created"`
}

// 订单详情传输模型
type WebOrderDetail struct {

	// 订单信息
	Id         uint    `json:"id"`
	Created    string  `json:"created"`
	UserName   string  `json:"username"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`

	// 地址信息
	Name            string `json:"name"`
	Mobile          string `json:"mobile"`
	PostalCode      int    `json:"postalCode"`
	Province        string `json:"province"`
	City            string `json:"city"`
	District        string `json:"district"`
	DetailedAddress string `json:"detailedAddress"`

	// 商品列表信息
	ProductItem []WebProductItem `json:"productItem"`
}

// 订单设置表单参数模型
type WebOrderSetInfo struct {
	Id              uint  `gorm:"primaryKey"        json:"id"`
	PaymentOvertime int64 `gorm:"payment_overtime"  json:"paymentOvertime"`
	ReceiveOvertime int64 `gorm:"receiveO_overtime" json:"receiveOvertime"`
	FinishOvertime  int64 `gorm:"finish_overtime"   json:"finishOvertime"`
}

type AppOrderListInfo struct {

	// 订单信息
	Id         uint    `json:"id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`

	// 商品列表信息
	ProductItem []AppProductItem `json:"productItem"`
}