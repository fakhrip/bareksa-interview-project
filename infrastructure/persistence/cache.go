package persistence

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func CreateRedisCache(redisPass string) *cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: redisPass,
		DB:       0,
	})

	mycache := cache.New(&cache.Options{
		Redis:      rdb,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return mycache
}
