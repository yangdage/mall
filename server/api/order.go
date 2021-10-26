package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var order service.Order
var orderSet service.OrderSet

// AppCreateOrder 创建订单
func AppCreateOrder(c *gin.Context) {
	var param models.AppOrderFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := order.AppCreate(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteOrder 删除订单
func WebDeleteOrder(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := order.WebDelete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateOrder 更新订单
func WebUpdateOrder(c *gin.Context) {
	var param models.WebOrderFormParam
	_ = c.ShouldBindJSON(&param)
	count := order.WebUpdate(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetOrderList 获取订单列表
func WebGetOrderList(c *gin.Context) {
	var page models.Page
	var param models.WebOrderFormParam
	_ = c.Bind(&page)
	_ = c.Bind(&param)
	orderList, row := order.WebGetList(page, param)
	response.SuccessPage("操作成功", orderList, row, c)
}

// AppGetOrderList 获取订单列表
func AppGetOrderList(c *gin.Context) {
	var param models.AppOrderFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	orderList := order.AppGetList(param)
	response.Success("操作成功", orderList, c)
}

// WebGetOrderDetail 获取订单详情
func WebGetOrderDetail(c *gin.Context) {
	var param models.PrimaryKey
	_ = c.Bind(&param)
	orderDetail := order.WebGetDetail(param.Id)
	response.Success("操作成功", orderDetail, c)
}

// WebSaveOrderSet 创建订单设置信息
func WebSaveOrderSet(c *gin.Context) {
	var param models.WebOrderSetFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := orderSet.WebSave(param)
	if count > 0 {
		response.Success("保存成功", count, c)
		return
	}
	response.Failed("保存失败", c)
}

// WebGetOrderSetInfo 获取订单设置信息
func WebGetOrderSetInfo(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		fmt.Println(err)
		return
	}
	info := orderSet.WebGetInfo(key.Id)
	response.Success("操作成功", info, c)
}
