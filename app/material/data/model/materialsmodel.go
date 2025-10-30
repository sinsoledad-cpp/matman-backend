package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MaterialsModel = (*customMaterialsModel)(nil)

type (
	// MaterialsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMaterialsModel.
	MaterialsModel interface {
		materialsModel
		withSession(session sqlx.Session) MaterialsModel
		CountAllByName(ctx context.Context, name string) (int64, error)
		FindAllByName(ctx context.Context, offset, limit int, name string) ([]*Materials, error)
		CountByFilters(ctx context.Context, name, materialType, supplierName string) (int64, error)
		FindAllByFilters(ctx context.Context, offset, limit int, name, materialType, supplierName string) ([]*Materials, error)
	}

	customMaterialsModel struct {
		*defaultMaterialsModel
	}
)

// NewMaterialsModel returns a model for the database table.
func NewMaterialsModel(conn sqlx.SqlConn) MaterialsModel {
	return &customMaterialsModel{
		defaultMaterialsModel: newMaterialsModel(conn),
	}
}

func (m *customMaterialsModel) withSession(session sqlx.Session) MaterialsModel {
	return NewMaterialsModel(sqlx.NewSqlConnFromSession(session))
}

// CountAllByName 根据名称模糊查询总数
func (m *customMaterialsModel) CountAllByName(ctx context.Context, name string) (int64, error) {
	var query string
	var args []interface{}

	likeName := "%" + name + "%"

	if name == "" {
		query = fmt.Sprintf("select count(*) from %s", m.table)
	} else {
		query = fmt.Sprintf("select count(*) from %s where `name` like ?", m.table)
		args = append(args, likeName)
	}

	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, args...)
	switch err {
	case nil:
		return count, nil
	case sqlx.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

// FindAllByName 根据名称模糊查询并分页
func (m *customMaterialsModel) FindAllByName(ctx context.Context, offset, limit int, name string) ([]*Materials, error) {
	var query string
	var args []interface{}

	likeName := "%" + name + "%"

	if name == "" {
		query = fmt.Sprintf("select %s from %s order by `id` asc limit ? offset ?", materialsRows, m.table)
		args = append(args, limit, offset)
	} else {
		query = fmt.Sprintf("select %s from %s where `name` like ? order by `id` asc limit ? offset ?", materialsRows, m.table)
		args = append(args, likeName, limit, offset)
	}

	var resp []*Materials
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// (新增) CountByFilters 根据动态条件查询总数
func (m *customMaterialsModel) CountByFilters(ctx context.Context, name, materialType, supplierName string) (int64, error) {
	var query strings.Builder
	var args []interface{}

	query.WriteString(fmt.Sprintf("select count(*) from %s where 1=1", m.table))

	if name != "" {
		query.WriteString(" AND `name` like ?")
		args = append(args, "%"+name+"%")
	}
	if materialType != "" {
		// 物料类型通常是精确匹配
		query.WriteString(" AND `material_type` = ?")
		args = append(args, materialType)
	}
	if supplierName != "" {
		// 供应商也使用模糊匹配
		query.WriteString(" AND `supplier_name` like ?")
		args = append(args, "%"+supplierName+"%")
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

// (新增) FindAllByFilters 根据动态条件分页查询
func (m *customMaterialsModel) FindAllByFilters(ctx context.Context, offset, limit int, name, materialType, supplierName string) ([]*Materials, error) {
	var query strings.Builder
	var args []interface{}

	query.WriteString(fmt.Sprintf("select %s from %s where 1=1", materialsRows, m.table))

	if name != "" {
		query.WriteString(" AND `name` like ?")
		args = append(args, "%"+name+"%")
	}
	if materialType != "" {
		query.WriteString(" AND `material_type` = ?")
		args = append(args, materialType)
	}
	if supplierName != "" {
		query.WriteString(" AND `supplier_name` like ?")
		args = append(args, "%"+supplierName+"%")
	}

	// 添加分页和排序
	query.WriteString(" order by `id` asc limit ? offset ?")
	args = append(args, limit, offset)

	var resp []*Materials
	err := m.conn.QueryRowsCtx(ctx, &resp, query.String(), args...)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
