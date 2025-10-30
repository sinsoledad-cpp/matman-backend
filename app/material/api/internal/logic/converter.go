package logic

import (
	"matman-backend/app/material/api/internal/types"
	"matman-backend/app/material/domain/entity"
	"time"
)

// Converter 负责 material 服务的 DTO 与 Entity 转换
type Converter struct {
	// (未来可以注入其他依赖，例如用于查询BOM子项名称)
}

// NewConverter 创建一个新的转换器实例
func NewConverter() *Converter {
	return &Converter{}
}

// ToMaterialInfoResponse 将 Material 实体转换为 DTO
func (c *Converter) ToMaterialInfoResponse(e *entity.Material) *types.MaterialInfo {
	if e == nil {
		return nil
	}

	return &types.MaterialInfo{
		Code:          e.Code,
		Name:          e.Name,
		MaterialType:  e.MaterialType,
		Spec:          e.Spec,
		Unit:          e.Unit,
		Price:         e.Price,
		StockQuantity: e.StockQuantity,
		SupplierName:  e.SupplierName,
		CreatedAt:     e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     e.UpdatedAt.Format(time.RFC3339),
	}
}

// ToMaterialInfoList (辅助函数) 转换实体列表
func (c *Converter) ToMaterialInfoList(entities []*entity.Material) []types.MaterialInfo {
	dtos := make([]types.MaterialInfo, len(entities))
	for i, e := range entities {
		dto := c.ToMaterialInfoResponse(e)
		if dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}

// ToBomEntryInfoResponse 将 BomEntry 实体转换为 DTO
func (c *Converter) ToBomEntryInfoResponse(e *entity.BomEntry) *types.BomEntryInfo {
	if e == nil {
		return nil
	}
	return &types.BomEntryInfo{
		ID:                 e.ID,
		ParentMaterialCode: e.ParentMaterialCode,
		ChildMaterialCode:  e.ChildMaterialCode,
		// (注意: ChildMaterialName 和 Spec 需要额外查询，这里暂时留空)
		// ChildMaterialName:  "",
		// ChildMaterialSpec:  "",
		Quantity:  e.Quantity,
		Status:    e.Status,
		CreatedAt: e.CreatedAt.Format(time.RFC3339),
		UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
	}
}

// ToBomEntryInfoList (辅助函数) 转换实体列表
func (c *Converter) ToBomEntryInfoList(entities []*entity.BomEntry) []types.BomEntryInfo {
	dtos := make([]types.BomEntryInfo, len(entities))
	for i, e := range entities {
		dto := c.ToBomEntryInfoResponse(e)
		if dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}

// (新增) ToBomListEntry 将 BomListDetails (领域读模型) 转换为 DTO
func (c *Converter) ToBomListEntry(e *entity.BomListDetails) *types.BomListEntry {
	if e == nil {
		return nil
	}
	return &types.BomListEntry{
		ID:                 e.ID,
		ParentMaterialCode: e.ParentMaterialCode,
		ParentMaterialName: e.ParentMaterialName,
		ChildMaterialCode:  e.ChildMaterialCode,
		ChildMaterialName:  e.ChildMaterialName,
		Quantity:           e.Quantity,
		Status:             e.Status,
		CreatedAt:          e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          e.UpdatedAt.Format(time.RFC3339),
	}
}

// (新增) ToBomListEntryList (辅助函数) 转换列表
func (c *Converter) ToBomListEntryList(entities []*entity.BomListDetails) []types.BomListEntry {
	dtos := make([]types.BomListEntry, len(entities))
	for i, e := range entities {
		dto := c.ToBomListEntry(e)
		if dto != nil {
			dtos[i] = *dto
		}
	}
	return dtos
}
