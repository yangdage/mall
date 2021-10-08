package models

// 类目列表传输模型
type CategoryList struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	ParentId uint   `json:"parentId"`
	Level    int    `json:"level"`
	Sort     int    `json:"sort"`
}

// 类目选项传输模型
type CategoryOption struct {
	Value    uint             `json:"value"`
	Label    string           `json:"label"`
	Children []CategoryOption `json:"children"`
}

