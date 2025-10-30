package entity

import "time"

// BomListDetails 是一个用于BOM列表查询的领域读模型
// 它组合了 BomEntry 以及父物料和子物料的详情。
type BomListDetails struct {
	ID                 int64
	ParentMaterialCode string
	ParentMaterialName string // (来自 JOIN)
	ChildMaterialCode  string
	ChildMaterialName  string // (来自 JOIN)
	Quantity           float64
	Status             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
