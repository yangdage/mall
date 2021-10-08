package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var product service.Product

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	var param models.ProductFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.Create(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// DeleteProduct 删除商品
func DeleteProduct(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// UpdateProduct 更新商品
func UpdateProduct(c *gin.Context) {
	var param models.ProductUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := product.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// GetProductInfo 获取商品信息
func GetProductInfo(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	productInfo := product.GetInfo(key.Id)
	response.Success("操作成功", productInfo, c)
}

// GetProductList 获取商品列表
func GetProductList(c *gin.Context) {
	var param models.ProductQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	productList, rows := product.GetList(param)
	response.SuccessPage("操作成功", productList, rows, c)
}
