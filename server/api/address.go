package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var address service.Address

// AppAddAddress 创建地址
func AppAddAddress(c *gin.Context) {
	var param models.AppAddressFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := address.AppAdd(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// AppDeleteAddress 删除地址
func AppDeleteAddress(c *gin.Context) {
	var key models.AppAddressDeleteParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := address.AppDelete(key.AddressId)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// AppUpdateAddress 更新地址
func AppUpdateAddress(c *gin.Context) {
	var param models.AppAddressUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := address.AppUpdate(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// AppGetAddressUpdateInfo 获取地址更新信息
func AppGetAddressUpdateInfo(c *gin.Context) {
	var key models.AppAddressInfoParam
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	info := address.AppGetUpdateInfo(key.AddressId)
	response.Success("操作成功", info, c)
}

// AppGetAddressList 获取地址列表
func AppGetAddressList(c *gin.Context) {
	var param models.AppAddressListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	addressList := address.AppGetList(param.UserId)
	response.Success("操作成功", addressList, c)
}
