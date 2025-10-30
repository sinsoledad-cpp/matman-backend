package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
		// (新增) 动态查询BOM总数
		CountByFilters(ctx context.Context, parentNameFilter string, statusFilter *int) (int64, error)
		// (新增) 动态查询BOM详情列表
		FindAllDetailsByFilters(ctx context.Context, offset, limit int, parentNameFilter string, statusFilter *int) ([]*BomListDetailsPO, error)
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

type BomListDetailsPO struct {
	Bom
	ParentName sql.NullString `db:"parent_name"`
	ChildName  sql.NullString `db:"child_name"`
}

// (新增) CountByFilters 动态查询BOM总数
func (m *customBomModel) CountByFilters(ctx context.Context, parentNameFilter string, statusFilter *int) (int64, error) {
	var query strings.Builder
	var args []interface{}

	// b: bom, p_mat: parent material
	query.WriteString(fmt.Sprintf(
		"SELECT count(b.id) FROM %s AS b ", m.table,
	))
	query.WriteString("LEFT JOIN `materials` AS p_mat ON b.parent_material_code = p_mat.code ")
	query.WriteString("WHERE 1=1")

	if parentNameFilter != "" {
		query.WriteString(" AND p_mat.name LIKE ?")
		args = append(args, "%"+parentNameFilter+"%")
	}
	if statusFilter != nil {
		query.WriteString(" AND b.status = ?")
		args = append(args, *statusFilter)
	}

	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query.String(), args...)
	switch err {
	case nil:
		return count, nil
	case sqlx.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// (新增) FindAllDetailsByFilters 动态查询BOM详情列表
func (m *customBomModel) FindAllDetailsByFilters(ctx context.Context, offset, limit int, parentNameFilter string, statusFilter *int) ([]*BomListDetailsPO, error) {
	var query strings.Builder
	var args []interface{}

	// b: bom, p_mat: parent material, c_mat: child material
	query.WriteString(fmt.Sprintf(
		"SELECT b.*, p_mat.name as parent_name, c_mat.name as child_name FROM %s AS b ", m.table,
	))
	query.WriteString("LEFT JOIN `materials` AS p_mat ON b.parent_material_code = p_mat.code ")
	query.WriteString("LEFT JOIN `materials` AS c_mat ON b.child_material_code = c_mat.code ")
	query.WriteString("WHERE 1=1")

	if parentNameFilter != "" {
		query.WriteString(" AND p_mat.name LIKE ?")
		args = append(args, "%"+parentNameFilter+"%")
	}
	if statusFilter != nil {
		query.WriteString(" AND b.status = ?")
		args = append(args, *statusFilter)
	}

	query.WriteString(" ORDER BY b.id ASC LIMIT ? OFFSET ?")
	args = append(args, limit, offset)

	var resp []*BomListDetailsPO
	err := m.conn.QueryRowsCtx(ctx, &resp, query.String(), args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return []*BomListDetailsPO{}, nil
	default:
		return nil, err
	}
}
