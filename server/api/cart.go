package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var cart service.Cart

func AppAddCart(c *gin.Context) {
	var param models.AppCartParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := cart.AppAdd(param.ProductId, param.UserId)
	if count > 0 {
		response.Success("添加成功", count, c)
		return
	}
	response.Failed("添加失败", c)
}

func AppDeleteCart(c *gin.Context) {
	var param models.AppCartParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := cart.AppDelete(param.ProductId, param.UserId)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

func AppClearCart(c *gin.Context) {
	var param models.AppCartParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := cart.AppClear(param.UserId)
	if count > 0 {
		response.Success("清除成功", count, c)
		return
	}
	response.Failed("清除失败", c)
}

func AppGetCartInfo(c *gin.Context) {
	var param models.AppCartParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	info := cart.AppGetInfo(param.UserId)
	if len(info.CartItem) == 0 {
		response.Success("购物车竟然是空的", info, c)
	}
	response.Success("操作成功", info, c)
}
