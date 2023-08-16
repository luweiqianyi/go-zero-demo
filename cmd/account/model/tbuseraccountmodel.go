package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserAccountModel = (*customTbUserAccountModel)(nil)

type (
	// TbUserAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserAccountModel.
	TbUserAccountModel interface {
		tbUserAccountModel
	}

	customTbUserAccountModel struct {
		*defaultTbUserAccountModel
	}
)

// NewTbUserAccountModel returns a model for the database table.
func NewTbUserAccountModel(conn sqlx.SqlConn) TbUserAccountModel {
	return &customTbUserAccountModel{
		defaultTbUserAccountModel: newTbUserAccountModel(conn),
	}
}
