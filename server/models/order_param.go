package models

// 订单表单参数模型
type WebOrderFormParam struct {
	Id             uint    `form:"id"             json:"id"`
	ProductItem    string  `form:"ProductItem"    json:"productItem"`
	TotalPrice     float64 `form:"totalPrice"     json:"totalPrice"`
	Status         string  `form:"status"         json:"status"`
	CourierName    string  `form:"courierName"    json:"courierName"`
	ShipmentNumber uint    `form:"shipmentNumber" json:"shipmentNumber"`
	UserId         string  `form:"userId"         json:"userId"`
	AdminId        uint    `form:"adminId"        json:"adminId"`
	AddressId      uint    `form:"addressId"      json:"addressId"`
}

// 订单设置表单参数模型
type WebOrderSetFormParam struct {
	Id              uint  `json:"id"`
	PaymentOvertime int64 `json:"paymentOvertime" binding:"required"`
	ReceiveOvertime int64 `json:"receiveOvertime" binding:"required"`
	FinishOvertime  int64 `json:"finishOvertime"  binding:"required"`
	AdminId         uint  `json:"adminId"`
}

// 订单表单参数模型
type AppOrderFormParam struct {
	UserId   string `form:"userId"    json:"userId"`
	NickName string `form:"nickName"  json:"nickName"`
	Status   string `form:"status"    json:"status"`
}
