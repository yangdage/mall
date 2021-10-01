package service

import (
	"mall.com/global"
	"mall.com/models"
)

type Statistics struct {

}

// GetStatisticsInfo 获取统计信息（商品数、订单量、交易金额）
func (s *Statistics) GetStatisticsInfo(id uint) models.StatisticsInfo {
	var info models.StatisticsInfo
	global.Db.Raw("SELECT COUNT(id) FROM `product` WHERE creator_id = ?", id).Scan(&info.GoodsCount)
	global.Db.Raw("SELECT COUNT(id) FROM `order` WHERE admin_id = ?", id).Scan(&info.OrderCount)
	global.Db.Raw("SELECT SUM(total_price) FROM `order` WHERE admin_id = ?", id).Scan(&info.Amount)
	return info
}