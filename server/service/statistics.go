package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
	"time"
)

type WebStatisticsService struct {

}

// GetOverviewInfo 后台管理前端，获取数据总览统计信息（商品数、订单量、交易金额）
func (s *WebStatisticsService) GetDataOverviewInfo() models.WebDataOverviewInfo {
	var info models.WebDataOverviewInfo
	global.Db.Raw("SELECT COUNT(id) FROM `product`").Scan(&info.GoodsCount)
	global.Db.Raw("SELECT COUNT(id) FROM `order`").Scan(&info.OrderCount)
	global.Db.Raw("SELECT SUM(total_price) FROM `order`").Scan(&info.Amount)
	return info
}

// GetTodayInfo 后台管理前端，获取今日订单数据统计信息
func (s *WebStatisticsService) GetTodayDataInfo() models.WebTodayOrderInfo {
	var todayInfo models.WebTodayOrderInfo
	today := time.Now().Format("2006-01-02")
	createdLike := today + "%"
	statusList := [5]string{"待付款","待发货","配送中","待收货","已完成"}

	for index, status := range statusList{
		selectSql := "SELECT COUNT(id) FROM `order` WHERE status = ? and created like ?"
		global.Db.Raw(selectSql, status, createdLike).Scan(&todayInfo.Data[index])
	}
	return todayInfo
}

// GetWeekInfo 后台管理前端，获取本周数据总览统计信息
func (s *WebStatisticsService) GetWeekDataInfo() models.WebWeekOrderInfo {
	var weekInfo models.WebWeekOrderInfo
	switch time.Now().Weekday() {
	case time.Monday:
		weekInfo = getWeekInfo(1)
	case time.Tuesday:
		weekInfo = getWeekInfo(2)
	case time.Wednesday:
		weekInfo = getWeekInfo(3)
	case time.Thursday:
		weekInfo = getWeekInfo(4)
	case time.Friday:
		weekInfo = getWeekInfo(5)
	case time.Saturday:
		weekInfo = getWeekInfo(6)
	case time.Sunday:
		weekInfo = getWeekInfo(7)
	default:
	}
	return weekInfo
}

func getWeekInfo(days int) models.WebWeekOrderInfo {
	var woi models.WebWeekOrderInfo
	for i,index := days-1, 0; i >= 0 ; i-- {
		var result []float64
		var amountMum float64
		nowTime := common.WeekTime(i) + "%"
		global.Db.Raw("SELECT COUNT(id) FROM `order` WHERE created like ?", nowTime).Scan(&woi.Orders[index])
		global.Db.Table("order").Where("created like ?", nowTime).Pluck("total_price",&result)
		for _,v := range result{
			amountMum += v
		}
		woi.Amount[index] = amountMum
		index++
	}
	return woi
}