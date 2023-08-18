package store

import (
	"context"
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
	})
	return gRedisStoreInstance
}

func Set(key string, value string, seconds int) error {
	return getStoreInstance().SetexCtx(context.Background(), key, value, seconds)
}

func Get(key string) (string, error) {
	return getStoreInstance().Get(key)
}
