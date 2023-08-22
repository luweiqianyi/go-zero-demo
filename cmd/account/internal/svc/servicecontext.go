package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/cmd/account-rpc/accountrpcservice"
	"go-zero-demo/cmd/account/internal/config"
	"go-zero-demo/cmd/account/model"
)

type ServiceContext struct {
	Config config.Config

	TbUserAccountModel model.TbUserAccountModel
	AccountRpcClient   accountrpcservice.AccountRpcService
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)

	return &ServiceContext{
		Config:             c,
		TbUserAccountModel: model.NewTbUserAccountModel(conn),
		AccountRpcClient:   accountrpcservice.NewAccountRpcService(zrpc.MustNewClient(c.AccountRpcConf)),
	}
}
