// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbUserAccountFieldNames          = builder.RawFieldNames(&TbUserAccount{})
	tbUserAccountRows                = strings.Join(tbUserAccountFieldNames, ",")
	tbUserAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(tbUserAccountFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbUserAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(tbUserAccountFieldNames, "`accountName`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbUserAccountModel interface {
		Insert(ctx context.Context, data *TbUserAccount) (sql.Result, error)
		FindOne(ctx context.Context, accountName string) (*TbUserAccount, error)
		Update(ctx context.Context, data *TbUserAccount) error
		Delete(ctx context.Context, accountName string) error
	}

	defaultTbUserAccountModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbUserAccount struct {
		AccountName string         `db:"accountName"`
		Password    sql.NullString `db:"password"`
	}
)

func newTbUserAccountModel(conn sqlx.SqlConn) *defaultTbUserAccountModel {
	return &defaultTbUserAccountModel{
		conn:  conn,
		table: "`TbUserAccount`",
	}
}

func (m *defaultTbUserAccountModel) Delete(ctx context.Context, accountName string) error {
	query := fmt.Sprintf("delete from %s where `accountName` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, accountName)
	return err
}

func (m *defaultTbUserAccountModel) FindOne(ctx context.Context, accountName string) (*TbUserAccount, error) {
	query := fmt.Sprintf("select %s from %s where `accountName` = ? limit 1", tbUserAccountRows, m.table)
	var resp TbUserAccount
	err := m.conn.QueryRowCtx(ctx, &resp, query, accountName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserAccountModel) Insert(ctx context.Context, data *TbUserAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, tbUserAccountRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AccountName, data.Password)
	return ret, err
}

func (m *defaultTbUserAccountModel) Update(ctx context.Context, data *TbUserAccount) error {
	query := fmt.Sprintf("update %s set %s where `accountName` = ?", m.table, tbUserAccountRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Password, data.AccountName)
	return err
}

func (m *defaultTbUserAccountModel) tableName() string {
	return m.table
}