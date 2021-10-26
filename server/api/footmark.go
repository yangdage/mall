package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var footmark service.Footmark

func AppAddFootmark(c *gin.Context) {
	var param models.AppFootmarkParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := footmark.AppAdd(param)
	if count > 0 {
		response.Success("添加成功", count, c)
		return
	}
	response.Failed("添加失败", c)
}

func AppDeleteFootmark(c *gin.Context) {
	var param models.AppCollectionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := footmark.AppDelete(param.UserId)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

func AppGetFootmarknList(c *gin.Context) {
	var param models.AppCollectionParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	info := footmark.AppGetList(param.UserId)
	response.Success("操作成功", info, c)
}

