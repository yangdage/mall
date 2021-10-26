package models

// 品牌表单参数模型
type WebBrandFormParam struct {
	Name string `json:"name" binding:"required"`
	Sort int    `json:"sort" binding:"required,gt=0"`
}

// 品牌更新参数模型
type WebBrandUpdateParam struct {
	Id   uint   `json:"id"   binding:"required,gt=0"`
	Name string `json:"name" binding:"required"`
	Sort int    `json:"sort" binding:"required,gt=0"`
}

// 品牌查询参数模型
type WebBrandQueryParam struct {
	Page Page
	Name string `form:"name" binding:"omitempty"`
}
