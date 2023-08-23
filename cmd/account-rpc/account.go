package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-demo/cmd/account-rpc/internal/config"
	"go-zero-demo/cmd/account-rpc/internal/server"
	"go-zero-demo/cmd/account-rpc/internal/svc"
	"go-zero-demo/cmd/account-rpc/pb"
	"go-zero-demo/pkg/store"
	"os"
	"path/filepath"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/account.yaml", "the config file")
var redisConfigFile = flag.String("redis", "etc/redis.yaml", "redis config file")

func main() {
	flag.Parse()

	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/account.yaml")
	flag.Set("f", fullPath)

	fullPath = filepath.Join(dir, "./etc/redis.yaml")
	flag.Set("redis", fullPath)
	var redisConf redis.RedisConf
	conf.MustLoad(*redisConfigFile, &redisConf)
	redisConf.PingTimeout = time.Second * 60
	logx.Infof("redis config: %#v", redisConf)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//conf := redis.RedisConf{
	//	Host:        "redis:6379",
	//	Type:        "node",
	//	Pass:        "",
	//	Tls:         false,
	//	NonBlock:    false,
	//	PingTimeout: time.Second,
	//}
	store.MustUseRedisStore(redisConf)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAccountRpcServiceServer(grpcServer, server.NewAccountRpcServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
