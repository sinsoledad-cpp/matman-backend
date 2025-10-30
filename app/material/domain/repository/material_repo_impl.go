package repository

import (
	"context"
	"database/sql"
	"errors"
	"matman-backend/app/material/data/model" // 引用 goctl model
	"matman-backend/app/material/domain/entity"

	"github.com/zeromicro/go-zero/core/logx"
)

// 确保 materialRepoImpl 实现了 MaterialRepository 接口
var _ MaterialRepository = (*materialRepoImpl)(nil)

type materialRepoImpl struct {
	materialsModel model.MaterialsModel // 依赖 goctl materials model
	bomModel       model.BomModel       // 依赖 goctl bom model (用于 IsMaterialInUse 检查)
}

// 构造函数：同时注入两个 model
func NewMaterialRepoImpl(materialsModel model.MaterialsModel, bomModel model.BomModel) MaterialRepository {
	return &materialRepoImpl{
		materialsModel: materialsModel,
		bomModel:       bomModel,
	}
}

// --- 辅助转换函数 (PO <-> Entity) ---

func toMaterialModel(e *entity.Material) *model.Materials {
	return &model.Materials{
		Id:            uint64(e.ID),
		Code:          e.Code,
		Name:          e.Name,
		MaterialType:  sql.NullString{String: e.MaterialType, Valid: e.MaterialType != ""},
		Spec:          sql.NullString{String: e.Spec, Valid: e.Spec != ""},
		Unit:          sql.NullString{String: e.Unit, Valid: e.Unit != ""},
		Price:         e.Price,
		StockQuantity: int64(e.StockQuantity),
		SupplierName:  sql.NullString{String: e.SupplierName, Valid: e.SupplierName != ""},
	}
}

func fromMaterialModel(m *model.Materials) *entity.Material {
	if m == nil {
		return nil
	}
	return &entity.Material{
		ID:            int64(m.Id),
		Code:          m.Code,
		Name:          m.Name,
		MaterialType:  m.MaterialType.String,
		Spec:          m.Spec.String,
		Unit:          m.Unit.String,
		Price:         m.Price,
		StockQuantity: int(m.StockQuantity),
		SupplierName:  m.SupplierName.String,
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
	}
}

// --- 接口实现 ---

func (r *materialRepoImpl) Create(ctx context.Context, mat *entity.Material) error {
	po := toMaterialModel(mat)
	res, err := r.materialsModel.Insert(ctx, po)
	if err != nil {
		logx.WithContext(ctx).Errorf("materialRepoImpl.Create error: %v", err)
		return err
	}
	lastId, _ := res.LastInsertId()
	mat.ID = lastId
	return nil
}

func (r *materialRepoImpl) Update(ctx context.Context, mat *entity.Material) error {
	poToUpdate, err := r.materialsModel.FindOneByCode(ctx, mat.Code)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return ErrNotFound
		}
		logx.WithContext(ctx).Errorf("materialRepoImpl.Update (FindOneByCode) error: %v", err)
		return err
	}

	// 更新字段
	poToUpdate.Name = mat.Name
	poToUpdate.MaterialType = sql.NullString{String: mat.MaterialType, Valid: mat.MaterialType != ""}
	poToUpdate.Spec = sql.NullString{String: mat.Spec, Valid: mat.Spec != ""}
	poToUpdate.Unit = sql.NullString{String: mat.Unit, Valid: mat.Unit != ""}
	poToUpdate.Price = mat.Price
	poToUpdate.StockQuantity = int64(mat.StockQuantity)
	poToUpdate.SupplierName = sql.NullString{String: mat.SupplierName, Valid: mat.SupplierName != ""}

	if err := r.materialsModel.Update(ctx, poToUpdate); err != nil {
		logx.WithContext(ctx).Errorf("materialRepoImpl.Update error: %v", err)
		return err
	}
	return nil
}

func (r *materialRepoImpl) DeleteByCode(ctx context.Context, code string) error {
	po, err := r.materialsModel.FindOneByCode(ctx, code)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return ErrNotFound
		}
		logx.WithContext(ctx).Errorf("materialRepoImpl.DeleteByCode (FindOneByCode) error: %v", err)
		return err
	}

	if err := r.materialsModel.Delete(ctx, po.Id); err != nil {
		logx.WithContext(ctx).Errorf("materialRepoImpl.DeleteByCode error: %v", err)
		return err
	}
	return nil
}

func (r *materialRepoImpl) FindByCode(ctx context.Context, code string) (*entity.Material, error) {
	po, err := r.materialsModel.FindOneByCode(ctx, code)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("materialRepoImpl.FindByCode error: %v", err)
		return nil, err
	}
	return fromMaterialModel(po), nil
}

func (r *materialRepoImpl) List(ctx context.Context, page, pageSize int, name, materialType, supplierName string) ([]*entity.Material, int64, error) {
	// (修改) 1. 调用新的 Count 方法
	total, err := r.materialsModel.CountByFilters(ctx, name, materialType, supplierName)
	if err != nil {
		logx.WithContext(ctx).Errorf("materialRepoImpl.List CountByFilters error: %v", err)
		return nil, 0, err
	}
	if total == 0 {
		return []*entity.Material{}, 0, nil
	}

	offset := (page - 1) * pageSize

	// (修改) 2. 调用新的 FindAll 方法
	pos, err := r.materialsModel.FindAllByFilters(ctx, offset, pageSize, name, materialType, supplierName)
	if err != nil {
		logx.WithContext(ctx).Errorf("materialRepoImpl.List FindAllByFilters error: %v", err)
		return nil, 0, err
	}

	// (不变) 3. 转换 PO -> Entity
	entities := make([]*entity.Material, len(pos))
	for i, po := range pos {
		entities[i] = fromMaterialModel(po)
	}
	return entities, total, nil
}

func (r *materialRepoImpl) IsMaterialInUse(ctx context.Context, code string) (bool, error) {
	return r.bomModel.CheckChildExists(ctx, code)
}
