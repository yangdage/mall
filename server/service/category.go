package service

import (
	"mall.com/common"
	"mall.com/global"
	"mall.com/models"
)

type Category struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"name"`
	ParentId uint   `gorm:"parent_id"`
	Level    int    `gorm:"level"`
	Sort     int    `gorm:"sort"`
	Created  string `gorm:"created"`
	Updated  string `gorm:"updated"`
}

// WebCreate 创建类目
func (c *Category) WebCreate(param models.WebCategoryFormParam) uint {
	var category Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = Category{
		Name:     param.Name,
		ParentId: param.ParentId,
		Level:    param.Level,
		Sort:     param.Sort,
		Created:  common.NowTime(),
	}
	if global.Db.Create(&category).RowsAffected > 0 {
		return category.Id
	}
	return 0
}

// WebDelete 删除类目
func (c *Category) WebDelete(id uint) int64 {
	var pid2, pid3 Category
	global.Db.Where("parent_id = ?", id).First(&pid2)
	global.Db.Where("parent_id = ?", pid2.Id).First(&pid3)
	return global.Db.Delete(&Category{}, []uint{id, pid2.Id, pid3.Id}).RowsAffected
}

// WebUpdate 更新类目
func (c *Category) WebUpdate(param models.WebCategoryUpdateParam) int64 {
	category := Category{
		Id:      param.Id,
		Name:    param.Name,
		Sort:    param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

// WebGetList 获取类目列表
func (c *Category) WebGetList(param models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {
	categoryList := make([]models.WebCategoryList, 0)
	query := &Category{
		Id:       param.Id,
		Name:     param.Name,
		Level:    param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &categoryList, &[]Category{})
	return categoryList, rows
}

// WebGetOption 获取类目选项
func (c *Category) WebGetOption() (option []models.WebCategoryOption) {
	selectList := make([]models.WebCategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

// 获取树形结构的选项
func getTreeOptions(id uint, cateList []models.WebCategoryList) (option []models.WebCategoryOption) {
	optionList := make([]models.WebCategoryOption, 0)
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2 || opt.Level == 3) {
			option := models.WebCategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 3 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}
