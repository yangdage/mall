package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var brand service.Brand

// WebCreateBrand 创建品牌
func WebCreateBrand(c *gin.Context) {
	var param models.WebBrandFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := brand.WebCreate(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteBrand 删除品牌
func WebDeleteBrand(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := brand.WebDelete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateBrand 更新品牌
func WebUpdateBrand(c *gin.Context) {
	var param models.WebBrandUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := brand.WebUpdate(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetBrandList 获取品牌列表
func WebGetBrandList(c *gin.Context) {
	var param models.WebBrandQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	brandList, rows := brand.WebGetList(param)
	response.SuccessPage("操作成功", brandList, rows, c)
}

// WebGetBrandOption 获取品牌选项
func WebGetBrandOption(c *gin.Context) {
	options := brand.WebGetOption()
	response.Success("操作成功", options, c)
}