package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var product service.Product

// WebCreateProduct 创建商品
func WebCreateProduct(c *gin.Context) {
	var param models.WebProductFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.WebCreate(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteProduct 删除商品
func WebDeleteProduct(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.WebDelete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateProduct 更新商品
func WebUpdateProduct(c *gin.Context) {
	var param models.WebProductUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.WebUpdate(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetProductInfo 获取商品信息
func WebGetProductInfo(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	productInfo := product.WebGetInfo(key.Id)
	response.Success("操作成功", productInfo, c)
}

// WebGetProductList 获取商品列表
func WebGetProductList(c *gin.Context) {
	var param models.WebProductQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	productList, rows := product.WebGetList(param)
	response.SuccessPage("操作成功", productList, rows, c)
}

// AppGetProductList 获取商品列表
func AppGetProductList(c *gin.Context) {
	productList := product.AppGetList()
	response.Success("操作成功", productList, c)
}

// AppGetProductDetail 获取商品详情
func AppGetProductDetail(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	detail := product.AppGetDetail(key.Id)
	response.Success("操作成功", detail, c)
}

