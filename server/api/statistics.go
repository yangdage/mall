package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var statistics service.Statistics

// WebGetStatisticsInfo 获取统计信息（商品数、订单量、交易金额）
func WebGetStatisticsInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.ShouldBind(&key)
	info := statistics.WebGetStatisticsInfo(key.Id)
	response.Success("操作成功", info, c)
}

// WebGetTodayOrderInfo 获取今日订单信息
func WebGetTodayOrderInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.ShouldBind(&key)
	info := statistics.WebGetTodayOrderInfo(key.Id)
	response.Success("操作成功", info, c)
}

// WebGetWeekInfo 获取每周订总览信息
func WebGetWeekInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.ShouldBind(&key)
	info := statistics.WebGetWeekInfo(key.Id)
	response.Success("操作成功", info, c)
}