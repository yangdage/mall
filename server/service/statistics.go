package service

import (
	"mall.com/global"
	"mall.com/models"
)

type Statistics struct {
	id           uint    `gorm:"primaryKey"`
	GoodsCount   int     `gorm:"goods_count"`
	OrderCount   int     `gorm:"order_count"`
	Amount       float64 `gorm:"amount"`
	VisitorCount int     `gorm:"visitor_count"`
	Created      string  `gorm:"created"`
	Updated      string  `gorm:"updated"`
}

func (s *Statistics) GetInfo() models.StatisticsInfo {
	var info models.StatisticsInfo
	global.Db.Table("statistics").Find(&info)
	return info
}
