package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BomModel = (*customBomModel)(nil)

type (
	// BomModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBomModel.
	BomModel interface {
		bomModel
		withSession(session sqlx.Session) BomModel
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
