package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var collection service.CollectionService

// AppAddCollection 微信小程序，添加收藏的商品
func AppAddCollection(c *gin.Context) {
	var param models.AppCollectionAddParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := collection.Add(param); count > 0 {
		response.Success("收藏成功", count, c)
		return
	}
	response.Failed("已收藏", c)
}

// AppDeleteCollection 微信小程序，删除收藏的商品
func AppDeleteCollection(c *gin.Context) {
	var param models.AppCollectionDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := collection.Delete(param); count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

// AppGetCollectionList 微信小程序，获取收藏的商品列表
func AppGetCollectionList(c *gin.Context) {
	var param models.AppCollectionQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	info := collection.GetList(param)
	response.Success("查询成功", info, c)
}

