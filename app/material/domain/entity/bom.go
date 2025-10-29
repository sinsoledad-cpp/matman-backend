package entity

import (
	"errors"
	"time"
)

// BomEntry 是 BOM 条目的领域实体
type BomEntry struct {
	ID                 int64
	ParentMaterialCode string
	ChildMaterialCode  string
	Quantity           float64
	Status             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// NewBomEntry 是创建 BOM 条目的工厂函数
// (已修正: 移除 DTO 依赖)
func NewBomEntry(parentCode, childCode string, quantity float64, status int) (*BomEntry, error) {
	if parentCode == "" || childCode == "" {
		return nil, errors.New("父物料和子物料编码不能为空")
	}
	if parentCode == childCode {
		return nil, errors.New("父物料和子物料不能是同一个")
	}
	if quantity <= 0 {
		return nil, errors.New("用量必须大于0")
	}
	// 确保 status 只有 0 或 1
	if status != 0 && status != 1 {
		status = 1 // 默认为 1 (生效)
	}

	return &BomEntry{
		ParentMaterialCode: parentCode,
		ChildMaterialCode:  childCode,
		Quantity:           quantity,
		Status:             status,
	}, nil
}
