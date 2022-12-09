package bootstrap

import (
	"fmt"

	"sample/pkg/redis"

	"sample/pkg/config"
)

func SetupRedis() {
	// 建立 Redis 连接
	redis.Connect(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
