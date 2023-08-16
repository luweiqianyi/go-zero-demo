package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserInfoModel = (*customTbUserInfoModel)(nil)

type (
	// TbUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserInfoModel.
	TbUserInfoModel interface {
		tbUserInfoModel
	}

	customTbUserInfoModel struct {
		*defaultTbUserInfoModel
	}
)

// NewTbUserInfoModel returns a model for the database table.
func NewTbUserInfoModel(conn sqlx.SqlConn) TbUserInfoModel {
	return &customTbUserInfoModel{
		defaultTbUserInfoModel: newTbUserInfoModel(conn),
	}
}
