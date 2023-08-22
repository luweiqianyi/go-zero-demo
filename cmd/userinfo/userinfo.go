package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go-zero-demo/cmd/userinfo/internal/config"
	"go-zero-demo/cmd/userinfo/internal/handler"
	"go-zero-demo/cmd/userinfo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/userinfo-api.yaml", "the config file")

func main() {
	flag.Parse()

	path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/userinfo-api.yaml")
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
