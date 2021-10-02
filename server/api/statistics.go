package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var statistics service.Statistics

// GetStatisticsInfo 获取统计信息（商品数、订单量、交易金额）
func GetStatisticsInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	info := statistics.GetStatisticsInfo(key.Id)
	response.Success("操作成功", info, c)
}

// GetTodayOrderInfo 获取今日订单信息
func GetTodayOrderInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	info := statistics.GetTodayOrderInfo(key.Id)
	response.Success("操作成功", info, c)
}

// GetWeekInfo 获取每周订总览信息
func GetWeekInfo(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	info := statistics.GetWeekInfo(key.Id)
	response.Success("操作成功", info, c)
}