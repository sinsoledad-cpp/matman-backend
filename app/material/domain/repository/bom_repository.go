package repository

import (
	"context"
	"errors"
	"matman-backend/app/material/domain/entity"
)

// ErrBomEntryExists 定义了 BOM 特有的仓储错误
var ErrBomEntryExists = errors.New("BOM entry already exists (unique constraint)")

// BomRepository 是 BOM 仓储的接口
type BomRepository interface {
	// CreateOrUpdate 尝试更新，如果 (parent, child) 唯一键不存在则创建
	CreateOrUpdate(ctx context.Context, bom *entity.BomEntry) (*entity.BomEntry, error)
	Delete(ctx context.Context, parentCode, childCode string) error
	FindByParentCode(ctx context.Context, parentCode string) ([]*entity.BomEntry, error)
	// FindEntry (用于内部检查唯一性)
	FindEntry(ctx context.Context, parentCode, childCode string) (*entity.BomEntry, error)
}
