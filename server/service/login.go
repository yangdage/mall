package service

import (
	"mall.com/global"
	"mall.com/models"
)

// WebLogin 用户登录信息验证
func (u *User) WebLogin(param models.WebUserParam) uint {
	var user User
	query := map[string]interface{}{
		"username": param.Username,
		"password": param.Password,
	}
	global.Db.Select("id").Where(query).First(&user)
	return user.Id
}
