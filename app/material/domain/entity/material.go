package entity

import (
	"errors"
	"time"
)

// Material 是物料领域的实体（Entity）
type Material struct {
	ID            int64
	Code          string
	Name          string
	MaterialType  string
	Spec          string
	Unit          string
	Price         int64 // 价格 (单位: 分)
	StockQuantity int
	SupplierName  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewMaterial 是创建物料的工厂函数，确保业务不变量
// (已修正: 移除 DTO 依赖，使用原生类型)
func NewMaterial(code, name, materialType, spec, unit, supplierName string, price int64, stockQuantity int) (*Material, error) {
	if code == "" || name == "" {
		return nil, errors.New("物料编码和名称不能为空")
	}
	if price < 0 {
		return nil, errors.New("价格不能为负数")
	}
	if stockQuantity < 0 {
		return nil, errors.New("库存不能为负数")
	}

	return &Material{
		// ID, CreatedAt, UpdatedAt 将由数据库生成
		Code:          code,
		Name:          name,
		MaterialType:  materialType,
		Spec:          spec,
		Unit:          unit,
		Price:         price,
		StockQuantity: stockQuantity,
		SupplierName:  supplierName,
	}, nil
}
