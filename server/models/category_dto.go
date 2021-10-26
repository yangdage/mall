package models

// 类目列表传输模型
type WebCategoryList struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId uint   `json:"parentId"`
	Level    int    `json:"level"`
	Sort     int    `json:"sort"`
}

// 类目选项传输模型
type WebCategoryOption struct {
	Value    uint                `json:"value"`
	Label    string              `json:"label"`
	Children []WebCategoryOption `json:"children"`
}
