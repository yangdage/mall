package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Product struct {
	Id           uint    `gorm:"primaryKey"`
	CategoryId   uint    `gorm:"category_id"`
	Kind         int     `gorm:"kind"`
	Title        string  `gorm:"title"`
	BrandId      uint    `gorm:"brand_id"`
	Name         string  `gorm:"name"`
	Price        float64 `gorm:"price"`
	Amount       int     `gorm:"amount"`
	Sales        int     `gorm:"sales"`
	ImageUrl     string  `gorm:"image_url"`
	SendAddress  string  `gorm:"send_address"`
	ParcelType   string  `gorm:"parcel_type"`
	SalesService string  `gorm:"sales_service"`
	CreatorId    uint    `gorm:"creator_id"`
	Status       int     `gorm:"status"`
	Created      string  `gorm:"created"`
	Updated      string  `gorm:"updated"`
}


// WebCreate 创建商品
func (p *Product) WebCreate(param models.WebProductFormParam) int64 {
	product := Product{
		CategoryId:     param.CategoryId,
		Kind:           param.Kind,
		Title:          param.Title,
		BrandId:        param.BrandId,
		Name:           param.Name,
		Price:          param.Price,
		Amount:         param.Amount,
		ImageUrl:       param.ImageUrl,
		SendAddress:    param.SendAddress,
		ParcelType:     param.ParcelType,
		SalesService:   param.SalesService,
		CreatorId:      param.CreatorId,
		Status:         param.Status,
		Created:        common.NowTime(),
	}
	return global.Db.Create(&product).RowsAffected
}

// WebDelete 删除商品
func (p *Product) WebDelete(id uint) int64 {
	return global.Db.Delete(&Product{}, id).RowsAffected
}

// WebUpdate 更新商品
func (p *Product) WebUpdate(param models.WebProductUpdateParam) int64 {
	product := Product{
		Id: 			param.Id,
		CategoryId:     param.CategoryId,
		Kind:           param.Kind,
		Title:          param.Title,
		BrandId:        param.BrandId,
		Name:           param.Name,
		Price:          param.Price,
		Amount:         param.Amount,
		ImageUrl:       param.ImageUrl,
		SendAddress:    param.SendAddress,
		ParcelType:     param.ParcelType,
		SalesService:   param.SalesService,
		CreatorId:      param.CreatorId,
		Status:         param.Status,
		Updated:        common.NowTime(),
	}
	return global.Db.Model(&product).Updates(product).RowsAffected
}

// WebGetInfo 获取商品信息
func (p *Product) WebGetInfo(id uint) models.WebProductInfo {
	var product models.WebProductInfo
	global.Db.Table("product").First(&product, id)
	return product
}

// WebGetList 获取商品列表
func (p *Product) WebGetList(param models.WebProductQueryParam) ([]models.WebProductList, int64) {
	query := &Product{
		Id:         param.Id,
		CategoryId: param.CategoryId,
		Title:      param.Title,
		BrandId:    param.BrandId,
		Status:     param.Status,
		CreatorId:  param.CreatorId,
	}
	productList := make([]models.WebProductList, 0)
	rows := common.RestPage(param.Page, "product", query, &productList, &[]Product{})
	return productList, rows
}

// AppGetList 获取商品列表
func (p *Product) AppGetList() []models.AppProductList {
	pList := make([]models.AppProductList, 0)
	global.Db.Table("product").Where("status = ?", 2).Find(&pList)
	return pList
}

// AppGetList 获取商品详情
func (p *Product) AppGetDetail(id uint) models.AppProductDetail {
	var detail models.AppProductDetail
	global.Db.Table("product").First(&detail, id)
	return detail
}

