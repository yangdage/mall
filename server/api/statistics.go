package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var statistics service.Statistics

// GetStatisticsInfo 获取统计信息
func GetStatisticsInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	info := statistics.GetInfo(key.Id)
	response.Success("操作成功", info, c)
}

// GetWeekInfo 获取本周信息
func GetWeekInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	weekOrders,weekAmount := statistics.GetWeekInfo(key.Id)
	weekInfo := map[string]interface{}{"orders": weekOrders, "amount": weekAmount}
	response.Success("操作成功", weekInfo, c)
}
