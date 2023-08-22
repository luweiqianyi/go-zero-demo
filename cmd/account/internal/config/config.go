package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	MySQL struct {
		DataSource string
	}

	Salt string

	AccountRpcConf zrpc.RpcClientConf
}
