package models

// 品牌列表传输模型
type BrandList struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Created string `json:"created"`
}

// 品牌选项传输模型
type BrandOption struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
