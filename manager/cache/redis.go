package cache

import (
	"com.pippishen/trans-proxy/manager"
	"fmt"
	"github.com/go-redis/redis"
)

func Redis() *redis.Client {
	redisCfg := manager.TP_CONFIG.Get("redis").(map[string]interface{})
	client := redis.NewClient(&redis.Options{
		Addr: redisCfg["addr"].(string),
		Password: redisCfg["password"].(string),
		DB: redisCfg["db"].(int),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("Redis connect ping failed, err: %s \n", err))
	} else {
		fmt.Printf("redis connect ping response: %s", pong)
		return client
	}
}