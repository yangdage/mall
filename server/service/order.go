package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
	"strconv"
	"strings"
)

type Order struct {
	Id             uint    `gorm:"primaryKey"`
	ProductItem    string  `gorm:"product_item"`
	TotalPrice     float64 `gorm:"total_price"`
	Status         string  `gorm:"status"`
	CourierName    string  `gorm:"courier_name"`
	ShipmentNumber uint    `gorm:"shipment_number"`
	AddressId      uint    `gorm:"address_id"`
	UserId         uint    `gorm:"user_id"`
	UserName       string  `gorm:"user_id"`
	AdminId        uint    `gorm:"admin_id"`
	Created        string  `gorm:"created"`
	Updated        string  `gorm:"updated"`
}

// Delete 删除订单
func (o *Order) Delete(id uint) int64 {
	return global.Db.Delete(&Order{}, id).RowsAffected
}

// Update 更新订单
func (o *Order) Update(param models.OrderParam) int64 {
	order := Order{
		Id:             param.Id,
		Status:         param.Status,
		CourierName:    param.CourierName,
		ShipmentNumber: param.ShipmentNumber,
		Updated:        common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// GetList 获取订单列表
func (o *Order) GetList(page models.Page, param models.OrderParam) ([]models.OrderList, int64) {
	orderList := make([]models.OrderList, 0)
	query := &Order{
		Id:      param.Id,
		Status:  param.Status,
		AdminId: param.AdminId,
	}
	rows := common.RestPage(page, "order", query, &orderList, &[]Order{})
	return orderList, rows
}

// GetDetail 获取订单详情
func (o *Order) GetDetail(orderId uint) (od models.OrderDetail) {
	var order Order
	var address models.Address
	var productItem []models.ProductItem
	// 查询订单信息
	global.Db.First(&order, orderId)
	// 查询地址信息
	global.Db.First(&address, order.AddressId)

	// 查询商品信息
	sList := strings.Split(order.ProductItem, ",")
	idList := make([]uint, 0)
	for _, pid := range sList {
		id, _ := strconv.Atoi(pid)
		if id != 0 {
			i := uint(id)
			idList = append(idList, i)
		}
	}
	global.Db.Table("product").Find(&productItem, idList)
	orderDetail := models.OrderDetail{
		Id:              order.Id,
		Created:         order.Created,
		Status:          order.Status,
		Name:            address.Name,
		Mobile:          address.Mobile,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderDetail
}
