package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var collection service.Collection

func AppAddCollection(c *gin.Context) {
	var param models.AppCollectionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := collection.AppAdd(param)
	if count > 0 {
		response.Success("收藏成功", count, c)
		return
	}
	response.Failed("已收藏", c)
}

func AppDeleteCollection(c *gin.Context) {
	var param models.AppCollectionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := collection.AppDelete(param.UserId)
	if count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

func AppGetCollectionList(c *gin.Context) {
	var param models.AppCollectionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	info := collection.AppGetList(param.UserId)
	response.Success("操作成功", info, c)
}

