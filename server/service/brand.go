package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Brand struct {
	Id      uint   `gorm:"primaryKey"`
	Name    string `gorm:"name"`
	Sort    int    `gorm:"sort"`
	Created string `gorm:"created"`
	Updated string `gorm:"updated"`
}

// WebCreate 创建品牌
func (b *Brand) WebCreate(param models.WebBrandFormParam) int64 {
	brand := Brand{
		Name:    param.Name,
		Sort:    param.Sort,
		Created: common.NowTime(),
	}
	return global.Db.Create(&brand).RowsAffected
}

// WebDelete 删除品牌
func (b *Brand) WebDelete(id uint) int64 {
	return global.Db.Delete(&Brand{}, id).RowsAffected
}

// WebUpdate 修改品牌
func (b *Brand) WebUpdate(param models.WebBrandUpdateParam) int64 {
	brand := Brand{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&brand).Updates(brand).RowsAffected
}

// WebGetList 获取品牌列表
func (b *Brand) WebGetList(param models.WebBrandQueryParam) ([]models.WebBrandList, int64) {
	brandList := make([]models.WebBrandList, 0)
	query := &Brand{ Name: param.Name }
	rows := common.RestPage(param.Page, "brand", query, &brandList, &[]Brand{})
	return brandList, rows
}

// WebGetOption 获取品牌选项
func (b *Brand) WebGetOption() []models.WebBrandOption {
	options := make([]models.WebBrandOption, 0)
	global.Db.Table("brand").Find(&options)
	return options
}
