package service

import (
	"fmt"
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
	UserId         string  `gorm:"user_id"`
	UserName       string  `gorm:"user_name"`
	AdminId        uint    `gorm:"admin_id"`
	Created        string  `gorm:"created"`
	Updated        string  `gorm:"updated"`
}

type OrderSet struct {
	Id              uint   `gorm:"id"`
	PaymentOvertime int64  `gorm:"payment_overtime"`
	ReceiveOvertime int64  `gorm:"receive_overtime"`
	FinishOvertime  int64  `gorm:"finish_overtime"`
	AdminId         uint   `gorm:"admin_id"`
	Created         string `gorm:"created"`
	Updated         string `gorm:"updated"`
}

var cart *Cart
var address *Address

// AppCreate 创建订单
func (o *Order) AppCreate(param models.AppOrderFormParam) int64 {
	info := cart.AppGetInfo(param.UserId)
	pids := make([]string, 0)
	for _, item := range info.CartItem {
		pids = append(pids, strconv.Itoa(int(item.Id)))
	}
	pidsItem := strings.Join(pids, ",")
	order := Order{
		ProductItem: pidsItem,
		TotalPrice:  info.TotalPrice,
		Status:      param.Status,
		AddressId:   address.AppGetId(param.UserId),
		UserId:      param.UserId,
		AdminId:     100030,
		UserName:    param.NickName,
		Created:     common.NowTime(),
	}
	return global.Db.Create(&order).RowsAffected
}

// WebDelete 删除订单
func (o *Order) WebDelete(id uint) int64 {
	return global.Db.Delete(&Order{}, id).RowsAffected
}

// WebUpdate 更新订单
func (o *Order) WebUpdate(param models.WebOrderFormParam) int64 {
	order := Order{
		Id:             param.Id,
		Status:         param.Status,
		CourierName:    param.CourierName,
		ShipmentNumber: param.ShipmentNumber,
		Updated:        common.NowTime(),
	}
	return global.Db.Model(&order).Updates(order).RowsAffected
}

// WebGetList 获取订单列表
func (o *Order) WebGetList(page models.Page, param models.WebOrderFormParam) ([]models.WebOrderList, int64) {
	orderList := make([]models.WebOrderList, 0)
	query := &Order{
		Id:      param.Id,
		Status:  param.Status,
		AdminId: param.AdminId,
	}
	rows := common.RestPage(page, "order", query, &orderList, &[]Order{})
	return orderList, rows
}

// AppGetList 获取订单列表
func (o *Order) AppGetList(param models.AppOrderFormParam) []models.AppOrderListInfo {

	// 查询满足特定条件的订单
	orderList := make([]Order, 0)
	query := &Order{
		UserId:  param.UserId,
		Status:  param.Status,
	}
	global.Db.Table("order").Where(query).Find(&orderList)
	orderInfoList := make([]models.AppOrderListInfo, 0)
	for _, order := range orderList {
		var orderInfo models.AppOrderListInfo
		productItem := make([]models.AppProductItem, 0)
		orderInfo.Id = order.Id
		orderInfo.Status = order.Status
		orderInfo.TotalPrice = order.TotalPrice

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
		orderInfo.ProductItem = productItem
		orderInfoList = append(orderInfoList, orderInfo)
	}
	return orderInfoList
}

// WebGetDetail 获取订单详情
func (o *Order) WebGetDetail(orderId uint) (od models.WebOrderDetail) {
	var order Order
	var address Address
	var productItem []models.WebProductItem

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
	orderDetail := models.WebOrderDetail{
		Id:              order.Id,
		Created:         order.Created,
		UserName:        order.UserName,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		Name:            address.Name,
		Mobile:          address.Mobile,
		PostalCode:      address.PostalCode,
		Province:        address.Province,
		City:            address.City,
		District:        address.District,
		DetailedAddress: address.DetailedAddress,
		ProductItem:     productItem,
	}
	return orderDetail
}

// WebSave 保存订单设置信息
func (o *OrderSet) WebSave(param models.WebOrderSetFormParam) int64 {
	orderSet := OrderSet{
		PaymentOvertime: param.PaymentOvertime,
		ReceiveOvertime: param.ReceiveOvertime,
		FinishOvertime:  param.FinishOvertime,
	}
	if row := global.Db.First(&OrderSet{}, param.Id).RowsAffected; row > 0 {
		orderSet.Id = param.Id
		orderSet.Updated = common.NowTime()
		fmt.Println(orderSet)
		return global.Db.Model(&orderSet).Updates(&orderSet).RowsAffected
	}
	orderSet.AdminId = param.AdminId
	orderSet.Created = common.NowTime()
	return global.Db.Create(&orderSet).RowsAffected
}

// WebGetInfo 获取订单设置信息
func (o *OrderSet) WebGetInfo(id uint) models.WebOrderSetInfo {
	var info models.WebOrderSetInfo
	global.Db.Table("order_set").Where("admin_id = ?", id).First(&info)
	return info
}
