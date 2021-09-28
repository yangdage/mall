package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/response"
	"mall.com/service"
)

var statistics service.Statistics

func GetStatisticsInfo(c *gin.Context) {
	info := statistics.GetInfo()
	response.Success("操作成功", info, c)
}
