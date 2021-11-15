package service

import (
	"mall.com/global"
	"mall.com/models"
)

type WebUserService struct {

}

// Login 用户登录信息验证
func (u *WebUserService) Login(param models.WebUserLoginParam) uint64 {
	var user models.User
	global.Db.Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return user.Id
}

