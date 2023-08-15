package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go-zero-demo/cmd/account/internal/config"
	"go-zero-demo/cmd/account/internal/handler"
	"go-zero-demo/cmd/account/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/account-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 设置配置文件路径为可执行文件所在路径+配置文件相对路径
	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/account-api.yaml")
	flag.Set("f", fullPath)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
