package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Avatar   string `gorm:"avatar"`
	Role     string `gorm:"role"`
	Status   int    `gorm:"status"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

func (u *User) WebUpdate(param models.WebUserParam) int64 {
	user := User{
		Id:       param.Id,
		Role:     param.Role,
		Status:   param.Status,
		Password: param.Password,
	}
	return global.Db.Model(&user).Updates(user).RowsAffected
}

func (u *User) WebDelete(id uint) int64 {
	return global.Db.Delete(&User{}, id).RowsAffected
}

func (u *User) WebGetList(page models.Page) (uList []models.WebUserList) {
	userList := make([]models.WebUserList, 0)
	common.RestPage(page, "user", &User{}, &userList, &User{})
	return userList
}
