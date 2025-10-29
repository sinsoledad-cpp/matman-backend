package repository

import (
	"context"
	"matman-backend/app/material/data/model" // 引用 goctl 生成的 model
	"matman-backend/app/material/domain/entity"
)

// ErrNotFound 从 data model 继承
var ErrNotFound = model.ErrNotFound

// MaterialRepository 是物料仓储的接口
type MaterialRepository interface {
	Create(ctx context.Context, mat *entity.Material) error
	Update(ctx context.Context, mat *entity.Material) error
	DeleteByCode(ctx context.Context, code string) error
	FindByCode(ctx context.Context, code string) (*entity.Material, error)
	List(ctx context.Context, page, pageSize int, name string) ([]*entity.Material, int64, error)

	// (这是一个业务查询，保留在这里是合理的，因为它查询的是 Material 的状态)
	IsMaterialInUse(ctx context.Context, code string) (bool, error)
}
