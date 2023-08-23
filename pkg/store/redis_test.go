package store

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
	"time"
)

func TestRedisUtil(t *testing.T) {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		PingTimeout: time.Second * 30,
	}
	MustUseRedisStore(conf)
	cli := getStoreInstance()
	success := cli.Ping()
	fmt.Println(success)

	if cli == nil {
		fmt.Println("end!")
	}
	err := Set("key", "TestRedisUtil", 3600)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
