package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BomModel = (*customBomModel)(nil)

type (
	// BomModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBomModel.
	BomModel interface {
		bomModel
		withSession(session sqlx.Session) BomModel
		FindAllByParentCode(ctx context.Context, parentCode string) ([]*Bom, error)
		CheckChildExists(ctx context.Context, childCode string) (bool, error)
	}

	customBomModel struct {
		*defaultBomModel
	}
)

// NewBomModel returns a model for the database table.
func NewBomModel(conn sqlx.SqlConn) BomModel {
	return &customBomModel{
		defaultBomModel: newBomModel(conn),
	}
}

func (m *customBomModel) withSession(session sqlx.Session) BomModel {
	return NewBomModel(sqlx.NewSqlConnFromSession(session))
}

// FindAllByParentCode 查询一个父物料的所有BOM条目
func (m *customBomModel) FindAllByParentCode(ctx context.Context, parentCode string) ([]*Bom, error) {
	query := fmt.Sprintf("select %s from %s where `parent_material_code` = ?", bomRows, m.table)
	var resp []*Bom
	err := m.conn.QueryRowsCtx(ctx, &resp, query, parentCode)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		// 未找到不应是错误，而是空列表
		return []*Bom{}, nil
	default:
		return nil, err
	}
}

// CheckChildExists 检查一个物料是否被用作子物料 (用于 IsMaterialInUse)
func (m *customBomModel) CheckChildExists(ctx context.Context, childCode string) (bool, error) {
	query := fmt.Sprintf("select exists(select 1 from %s where `child_material_code` = ? limit 1)", m.table)
	var exists bool
	err := m.conn.QueryRowCtx(ctx, &exists, query, childCode)
	if err != nil {
		return false, err
	}
	return exists, nil
}
