package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ MaterialsModel = (*customMaterialsModel)(nil)

type (
	// MaterialsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMaterialsModel.
	MaterialsModel interface {
		materialsModel
		withSession(session sqlx.Session) MaterialsModel
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
