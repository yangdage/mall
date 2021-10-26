package api

import (
	"github.com/gin-gonic/gin"
	"mall.com/models"
	"mall.com/response"
	"mall.com/service"
)

var user service.User

func WebUpdateUser(c *gin.Context) {
	var param models.WebUserParam
	_ = c.ShouldBindJSON(&param)
	count := user.WebUpdate(param)
	if count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

func WebDeleteUser(c *gin.Context) {
	var key models.PrimaryKey
	_ = c.Bind(&key)
	count := user.WebDelete(key.Id)
	if count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

func WebGetUserList(c *gin.Context) {
	var page models.Page
	_ = c.Bind(&page)
	userList := user.WebGetList(page)
	response.Success("操作成功", userList, c)
}
