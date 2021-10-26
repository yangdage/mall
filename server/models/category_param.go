package models

// 类目表单参数模型
type WebCategoryFormParam struct {
	Name     string `json:"name"     binding:"required"`
	ParentId uint   `json:"parentId" binding:"required,gt=0"`
	Level    int    `json:"level"    binding:"required,oneof=1 2 3"`
	Sort     int    `json:"sort"     binding:"required,gt=0"`
}

// 类目更新参数模型
type WebCategoryUpdateParam struct {
	Id   uint   `json:"id"       binding:"required,gt=0"`
	Name string `json:"name"     binding:"required"`
	Sort int    `json:"sort"     binding:"required,gt=0"`
}

// 类目查询参数模型
type WebCategoryQueryParam struct {
	Page     Page
	Id       uint   `form:"id"       binding:"omitempty,gt=0"`
	Name     string `form:"name"     binding:"omitempty"`
	ParentId uint   `form:"parentId" binding:"omitempty,gt=0"`
	Level    int    `form:"level"    binding:"omitempty,oneof=1 2 3"`
}
