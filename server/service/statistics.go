package service

import (
	"mall.com/global"
	"mall.com/models"
	"strconv"
	"time"
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

func (s *Statistics) GetInfo(id uint) models.StatisticsInfo {
	var info models.StatisticsInfo
	global.Db.Table("statistics").Where("admin_id = ?", id).Find(&info)
	return info
}

func (s *Statistics) GetWeekInfo(id uint) (weekOrders, weekAmount [7]int) {

	nowDay := [7]string{"monday","tuesday","wednesday","thursday","friday","saturday","sunday"}
	switch time.Now().Weekday() {
	case time.Monday:
		weekOrders,weekAmount = weekData(id, nowDay, 1)
	case time.Tuesday:
		weekOrders,weekAmount = weekData(id, nowDay, 2)
	case time.Wednesday:
		weekOrders,weekAmount = weekData(id, nowDay, 3)
	case time.Thursday:
		weekOrders,weekAmount = weekData(id, nowDay, 4)
	case time.Friday:
		weekOrders,weekAmount = weekData(id, nowDay, 5)
	case time.Saturday:
		weekOrders,weekAmount = weekData(id, nowDay, 6)
	case time.Sunday:
		weekOrders,weekAmount = weekData(id, nowDay, 7)
	default:
	}
	return weekOrders,weekAmount
}

func weekData(id uint, nowDay [7]string, times int) (weekOrders, weekAmount [7]int) {
	aid := strconv.Itoa(int(id))
	for i := 0; i < times; i++ {
		key := "statistics:" + aid + ":" + nowDay[i]
		value := global.RDb.HGetAll(ctx, key).Val()
		weekOrders[i], _ = strconv.Atoi(value["orders"])
		weekAmount[i], _ = strconv.Atoi(value["amount"])
	}
	return weekOrders, weekAmount
}

