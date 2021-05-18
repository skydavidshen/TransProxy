package cache

import (
	"TransProxy/manager"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: manager.TP_SERVER_CONFIG.Redis.Addr,
		Password: manager.TP_SERVER_CONFIG.Redis.Password,
		DB: manager.TP_SERVER_CONFIG.Redis.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Redis connect ping failed, err: %s \n", err))
	} else {
		fmt.Printf("redis connect ping response: %s", pong)
		return client
	}
}

func Close()  {
	err := manager.TP_CACHE_REDIS.Close()
	if err != nil {
		manager.TP_LOG.Error("redis client close failed, err:",
			zap.String("err", err.Error()))
	}
}
