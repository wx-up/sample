package redis

import (
	"context"
	"sync"
	"time"

	"sample/pkg/logger"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

var once sync.Once

func Connect(address string, username string, passwords string, db int) {
	once.Do(func() {
		Client = redis.NewClient(&redis.Options{
			Addr:     address,
			Username: username,
			Password: passwords,
			DB:       db,
		})
		if err := Ping(); err != nil {
			panic(err)
		}
	})
}

func Ping() error {
	_, err := Client.Ping(context.Background()).Result()
	return err
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func Set(key string, value any, expiration time.Duration) bool {
	return SetWithContext(context.Background(), key, value, expiration)
}

// Get 获取 key 对应的 value
func Get(key string) string {
	return GetWithContext(context.Background(), key)
}

// Has 判断一个 key 是否存在
func Has(key string) bool {
	return HasWithContext(context.Background(), key)
}

// Del 删除存储在 redis 里的数据，支持多个 key 传参
func Del(keys ...string) bool {
	return DelWithContext(context.Background(), keys...)
}

// Incr key 值 +1
func Incr(key string) bool {
	return IncrWithContext(context.Background(), key)
}

// IncrBy key 值 +val
func IncrBy(key string, val int64) bool {
	return IncrByWithContext(context.Background(), key, val)
}

// Decr key 值 -1
func Decr(key string) bool {
	return DecrWithContext(context.Background(), key)
}

// DecrBy key 值  -val
func DecrBy(key string, val int64) bool {
	return DecrByWithContext(context.Background(), key, val)
}

func SetWithContext(ctx context.Context, key string, value any, expiration time.Duration) bool {
	if err := Client.Set(ctx, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "SetWithContext", err.Error())
		return false
	}
	return true
}

func GetWithContext(ctx context.Context, key string) string {
	result, err := Client.Get(ctx, key).Result()
	if err != nil {
		// key 不存在的错误，不打印日志
		if err != redis.Nil {
			logger.ErrorString("Redis", "GetWithContext", err.Error())
		}
		return ""
	}
	return result
}

func HasWithContext(ctx context.Context, key string) bool {
	_, err := Client.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "HasWithContext", err.Error())
		}
		return false
	}
	return true
}

func DelWithContext(ctx context.Context, keys ...string) bool {
	if err := Client.Del(ctx, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "DelWithContext", err.Error())
		return false
	}
	return true
}

func IncrWithContext(ctx context.Context, key string) bool {
	if err := Client.Incr(ctx, key).Err(); err != nil {
		logger.ErrorString("Redis", "IncrWithContext", err.Error())
		return false
	}
	return true
}

func IncrByWithContext(ctx context.Context, key string, val int64) bool {
	if err := Client.IncrBy(ctx, key, val).Err(); err != nil {
		logger.ErrorString("Redis", "IncrByWithContext", err.Error())
		return false
	}
	return true
}

func DecrWithContext(ctx context.Context, key string) bool {
	if err := Client.Decr(ctx, key).Err(); err != nil {
		logger.ErrorString("Redis", "DecrWithContext", err.Error())
		return false
	}
	return true
}

func DecrByWithContext(ctx context.Context, key string, val int64) bool {
	if err := Client.DecrBy(ctx, key, val).Err(); err != nil {
		logger.ErrorString("Redis", "DecrByWithContext", err.Error())
		return false
	}
	return true
}
