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