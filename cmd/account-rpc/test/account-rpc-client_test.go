package test

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/cmd/account-rpc/accountrpcservice"
	"go-zero-demo/cmd/account-rpc/pb"
	"testing"
)

func TestRpcClient(t *testing.T) {
	// 在docker-compose.yml中将宿主机的8003映射到容器的8003端口,所以，在宿主机上使用地址
	// 127.0.0.1:8003是可以访问我们Docker中启动的account-rpc容器的
	c := zrpc.RpcClientConf{
		Target: "127.0.0.1:8003",
	}
	client := accountrpcservice.NewAccountRpcService(zrpc.MustNewClient(c))
	resp, err := client.ValidateToken(context.Background(), &pb.TokenValidateReq{})
	fmt.Printf("resp: %#v, err: %v\n", resp, err)
}
