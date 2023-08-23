package store

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
)

var gConf redis.RedisConf
var redisStoreOnce sync.Once
var gRedisStoreInstance *redis.Redis

func MustUseRedisStore(conf redis.RedisConf) {
	gConf = conf
}

func getStoreInstance() *redis.Redis {
	redisStoreOnce.Do(func() {
		gRedisStoreInstance = redis.MustNewRedis(gConf)
		if gRedisStoreInstance != nil {
			logx.Infof("connect to %#v success", gConf)
		} else {
			logx.Errorf("redis cli nil")
		}
	})
	return gRedisStoreInstance
}

func Set(key string, value string, seconds int) error {
	redisUtil := getStoreInstance()
	if redisUtil == nil {
		return fmt.Errorf("redis cli required")
	}
	return redisUtil.SetexCtx(context.Background(), key, value, seconds)
}

func Get(key string) (string, error) {
	redisUtil := getStoreInstance()

	if redisUtil == nil {
		return "", fmt.Errorf("redis cli required")
	}

	return redisUtil.Get(key)
}
