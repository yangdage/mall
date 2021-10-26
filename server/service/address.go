package service

import (
	"fmt"
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Address struct {
	Id              uint   `gorm:"primaryKey"`
	UserId          string `gorm:"user_id"`
	Name            string `gorm:"name"`
	Mobile          string `gorm:"mobile"`
	PostalCode      int    `gorm:"postal_code"`
	Province        string `gorm:"province"`
	City            string `gorm:"city"`
	District        string `gorm:"district"`
	DetailedAddress string `gorm:"detailed_address"`
	IsDefault       int    `gorm:"is_default"`
	Created         string `gorm:"created"`
	Updated         string `gorm:"updated"`
}

// AppAdd 新建地址
func (a *Address) AppAdd(param models.AppAddressFormParam) int64 {
	address := Address{
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Created:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Create(&address).RowsAffected
		}
		return global.Db.Create(&address).RowsAffected
	}
	return global.Db.Create(&address).RowsAffected
}

// AppDelete 删除地址
func (a *Address) AppDelete(id uint) int64 {
	return global.Db.Delete(&Address{}, id).RowsAffected
}

// AppUpdate 更新地址
func (a *Address) AppUpdate(param models.AppAddressUpdateParam) int64 {
	address := Address{
		Id:              param.Id,
		UserId:          param.UserId,
		Name:            param.Name,
		Mobile:          param.Mobile,
		PostalCode:      param.PostalCode,
		Province:        param.Province,
		City:            param.City,
		District:        param.District,
		DetailedAddress: param.DetailedAddress,
		IsDefault:       param.IsDefault,
		Updated:         common.NowTime(),
	}
	if address.IsDefault == 1 {
		var addressId uint
		row := global.Db.Table("address").Select("id").
			Where("is_default = ? and user_id = ?", 1, address.UserId).Take(&addressId).RowsAffected
		fmt.Println(addressId)
		if row > 0 {
			global.Db.Table("address").Where("id = ?", addressId).
				Update("is_default", 2)
			return global.Db.Updates(&address).RowsAffected
		}
		return global.Db.Updates(&address).RowsAffected
	}
	return global.Db.Updates(&address).RowsAffected
}

// AppGetId 获取地址Id
func (a *Address) AppGetId(uid string) uint {
	var id uint
	global.Db.Table("address").Select("id").
		      Where("is_default = ? and user_id = ?", 1, uid).Take(&id)
	return id
}

// AppGetUpdateInfo 获取更新信息
func (a *Address) AppGetUpdateInfo(id uint) models.AppAddressUpdateInfo {
	var updateInfo models.AppAddressUpdateInfo
	global.Db.Table("address").First(&updateInfo, id)
	return updateInfo
}

// AppGetList 获取地址列表
func (a *Address) AppGetList(uid string) []models.AppAddressList {
	aList := make([]models.AppAddressList, 0)
	global.Db.Table("address").Where("user_id = ?", uid).Order("is_default").Find(&aList)
	return aList
}
