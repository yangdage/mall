package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var category service.Category

// WebCreateCategory 创建类目
func WebCreateCategory(c *gin.Context) {
	var param models.WebCategoryFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		fmt.Println(err)
		return
	}
	count := category.WebCreate(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteCategory 删除类目
func WebDeleteCategory(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := category.WebDelete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateCategory 更新类目
func WebUpdateCategory(c *gin.Context) {
	var param models.WebCategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := category.WebUpdate(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetCategoryList 获取类目列表
func WebGetCategoryList(c *gin.Context) {
	var param models.WebCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	cateList, rows := category.WebGetList(param)
	response.SuccessPage("操作成功", cateList, rows, c)
}

// WebGetCategoryOption 获取类目选项
func WebGetCategoryOption(c *gin.Context) {
	option := category.WebGetOption()
	response.Success("操作成功", option, c)
}
