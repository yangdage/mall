package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var category service.Category

// CreateCategory 创建类目
func CreateCategory(c *gin.Context) {
	var param models.CategoryFormParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		fmt.Println(err)
		return
	}
	count := category.Create(param)
	if count > 0 {
		response.Success("创建成功", count, c)
		return
	}
	response.Failed("创建失败", c)
}

// DeleteCategory 删除类目
func DeleteCategory(c *gin.Context) {
	var key models.PrimaryKey
	if err := c.ShouldBind(&key); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := category.Delete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// UpdateCategory 更新类目
func UpdateCategory(c *gin.Context) {
	var param models.CategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	count := category.Update(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// GetCategoryList 获取类目列表
func GetCategoryList(c *gin.Context) {
	var param models.CategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("参数无效", c)
		return
	}
	cateList, rows := category.GetList(param)
	response.SuccessPage("操作成功", cateList, rows, c)
}

// GetCategoryOption 获取类目选项
func GetCategoryOption(c *gin.Context) {
	option := category.GetOption()
	response.Success("操作成功", option, c)
}
