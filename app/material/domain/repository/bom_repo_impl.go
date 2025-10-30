package repository

import (
	"context"
	"errors"
	"matman-backend/app/material/data/model" // 引用 goctl model
	"matman-backend/app/material/domain/entity"
	// 移除了 "repo" 别名导入

	"github.com/zeromicro/go-zero/core/logx"
)

// 确保 bomRepoImpl 实现了 BomRepository 接口
var _ BomRepository = (*bomRepoImpl)(nil)

type bomRepoImpl struct {
	bomModel model.BomModel // 依赖 goctl bom model
}

func NewBomRepoImpl(bomModel model.BomModel) BomRepository {
	return &bomRepoImpl{
		bomModel: bomModel,
	}
}

// --- 辅助转换函数 (PO <-> Entity) ---

func toBomModel(e *entity.BomEntry) *model.Bom {
	return &model.Bom{
		Id:                 uint64(e.ID),
		ParentMaterialCode: e.ParentMaterialCode,
		ChildMaterialCode:  e.ChildMaterialCode,
		Quantity:           e.Quantity,
		Status:             int64(e.Status),
	}
}

func fromBomModel(m *model.Bom) *entity.BomEntry {
	if m == nil {
		return nil
	}
	return &entity.BomEntry{
		ID:                 int64(m.Id),
		ParentMaterialCode: m.ParentMaterialCode,
		ChildMaterialCode:  m.ChildMaterialCode,
		Quantity:           m.Quantity,
		Status:             int(m.Status),
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}
func fromBomListDetailsPO(m *model.BomListDetailsPO) *entity.BomListDetails {
	if m == nil {
		return nil
	}
	return &entity.BomListDetails{
		ID:                 int64(m.Id),
		ParentMaterialCode: m.ParentMaterialCode,
		ParentMaterialName: m.ParentName.String,
		ChildMaterialCode:  m.ChildMaterialCode,
		ChildMaterialName:  m.ChildName.String,
		Quantity:           m.Quantity,
		Status:             int(m.Status),
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}

// --- 接口实现 ---

func (r *bomRepoImpl) FindEntry(ctx context.Context, parentCode, childCode string) (*entity.BomEntry, error) {
	po, err := r.bomModel.FindOneByParentMaterialCodeChildMaterialCode(ctx, parentCode, childCode)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("bomRepoImpl.FindEntry error: %v", err)
		return nil, err
	}
	return fromBomModel(po), nil
}

func (r *bomRepoImpl) CreateOrUpdate(ctx context.Context, bom *entity.BomEntry) (*entity.BomEntry, error) {
	po, err := r.bomModel.FindOneByParentMaterialCodeChildMaterialCode(ctx, bom.ParentMaterialCode, bom.ChildMaterialCode)

	if err == nil {
		// 找到了 -> 更新
		po.Quantity = bom.Quantity
		po.Status = int64(bom.Status)
		if err := r.bomModel.Update(ctx, po); err != nil {
			logx.WithContext(ctx).Errorf("bomRepoImpl.CreateOrUpdate (Update) error: %v", err)
			return nil, err
		}
		return fromBomModel(po), nil

	} else if errors.Is(err, model.ErrNotFound) {
		// 没找到 -> 创建
		poToInsert := toBomModel(bom)
		res, err := r.bomModel.Insert(ctx, poToInsert)
		if err != nil {
			logx.WithContext(ctx).Errorf("bomRepoImpl.CreateOrUpdate (Insert) error: %v", err)
			return nil, err
		}
		lastId, _ := res.LastInsertId()
		bom.ID = lastId
		return bom, nil

	} else {
		// 其他错误
		logx.WithContext(ctx).Errorf("bomRepoImpl.CreateOrUpdate (Find) error: %v", err)
		return nil, err
	}
}

func (r *bomRepoImpl) Delete(ctx context.Context, parentCode, childCode string) error {
	po, err := r.bomModel.FindOneByParentMaterialCodeChildMaterialCode(ctx, parentCode, childCode)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return ErrNotFound
		}
		logx.WithContext(ctx).Errorf("bomRepoImpl.Delete (Find) error: %v", err)
		return err
	}

	if err := r.bomModel.Delete(ctx, po.Id); err != nil {
		logx.WithContext(ctx).Errorf("bomRepoImpl.Delete error: %v", err)
		return err
	}
	return nil
}

func (r *bomRepoImpl) FindByParentCode(ctx context.Context, parentCode string) ([]*entity.BomEntry, error) {
	pos, err := r.bomModel.FindAllByParentCode(ctx, parentCode)
	if err != nil {
		// 如果是未找到，返回空切片
		if errors.Is(err, model.ErrNotFound) {
			return []*entity.BomEntry{}, nil
		}
		logx.WithContext(ctx).Errorf("bomRepoImpl.FindByParentCode error: %v", err)
		return nil, err
	}

	entities := make([]*entity.BomEntry, len(pos))
	for i, po := range pos {
		entities[i] = fromBomModel(po)
	}
	return entities, nil
}

// (新增) ListDetails 实现
func (r *bomRepoImpl) ListDetails(ctx context.Context, page, pageSize int, parentNameFilter string, statusFilter *int) ([]*entity.BomListDetails, int64, error) {
	// 1. 获取总数
	total, err := r.bomModel.CountByFilters(ctx, parentNameFilter, statusFilter)
	if err != nil {
		logx.WithContext(ctx).Errorf("bomRepoImpl.ListDetails CountByFilters error: %v", err)
		return nil, 0, err
	}
	if total == 0 {
		return []*entity.BomListDetails{}, 0, nil
	}

	// 2. 计算 offset
	offset := (page - 1) * pageSize

	// 3. 查询分页数据
	pos, err := r.bomModel.FindAllDetailsByFilters(ctx, offset, pageSize, parentNameFilter, statusFilter)
	if err != nil {
		logx.WithContext(ctx).Errorf("bomRepoImpl.ListDetails FindAllDetailsByFilters error: %v", err)
		return nil, 0, err
	}

	// 4. 转换 PO -> Entity
	entities := make([]*entity.BomListDetails, len(pos))
	for i, po := range pos {
		entities[i] = fromBomListDetailsPO(po)
	}

	return entities, total, nil
}
