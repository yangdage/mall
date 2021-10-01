package models

type OrderParam struct {
	Id             uint    `form:"id"             json:"id"`
	ProductItem    string  `form:"ProductItem"    json:"productItem"`
	TotalPrice     float64 `form:"totalPrice"     json:"totalPrice"`
	Status         string  `form:"status"         json:"status"`
	CourierName    string  `form:"courierName"    json:"courierName"`
	ShipmentNumber uint    `form:"shipmentNumber" json:"shipmentNumber"`
	UserId         uint    `form:"userId"         json:"userId"`
	AdminId        uint    `form:"adminId"        json:"adminId"`
	AddressId      uint    `form:"addressId"      json:"addressId"`
}
