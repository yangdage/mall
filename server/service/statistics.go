package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
	"time"
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

// GetTodayOrderInfo 获取今日订单信息
func (s *Statistics) GetTodayOrderInfo(id uint) models.TodayOrderInfo {
	var toi models.TodayOrderInfo
	today := time.Now().Format("2006-01-02")
	createdLike := today + "%"
	statusList := [6]string{"待付款","待发货","已发货","待收货","待评价","已完成"}

	for index, status := range statusList{
		sql := "SELECT COUNT(id) FROM `order` WHERE admin_id = ? and status = ? and created like ?"
		global.Db.Raw(sql, id, status, createdLike).Scan(&toi.Data[index])
	}
	return toi
}

// GetWeekInfo 获取本周订单总览信息
func (s *Statistics) GetWeekInfo(id uint) models.WeekOrderInfo {
	var woi models.WeekOrderInfo
	switch time.Now().Weekday() {
	case time.Monday:
		woi = weekInfo(id, 1)
	case time.Tuesday:
		woi = weekInfo(id, 2)
	case time.Wednesday:
		woi = weekInfo(id, 3)
	case time.Thursday:
		woi = weekInfo(id, 4)
	case time.Friday:
		woi = weekInfo(id, 5)
	case time.Saturday:
		woi = weekInfo(id, 6)
	case time.Sunday:
		woi = weekInfo(id, 7)
	default:
	}
	return woi
}

func weekInfo(id uint, days int) models.WeekOrderInfo {
	var woi models.WeekOrderInfo
	for i,index := days-1, 0; i >= 0 ; i-- {
		var result []float64
		var amountMum float64
		nowTime := common.WeekTime(i) + "%"
		ordersSql := "SELECT COUNT(id) FROM `order` WHERE admin_id = ? and created like ?"
		amountSql := "admin_id = ? and created like ?"
		global.Db.Raw(ordersSql, id, nowTime).Scan(&woi.Orders[index])
		global.Db.Table("order").Where(amountSql, id, nowTime).Pluck("total_price",&result)
		for _,v := range result{
			amountMum += v
		}
		woi.Amount[index] = amountMum
		index++
	}
	return woi
}