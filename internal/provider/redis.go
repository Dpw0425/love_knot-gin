package provider

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"love_knot/internal/config"
	"love_knot/pkg/logger"
)

func NewRedisClient(conf *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        conf.Redis.Host + ":" + conf.Redis.Port,
		Password:    conf.Redis.Auth,
		DB:          conf.Redis.Database,
		ReadTimeout: 0,
	})

	if _, err := client.WithContext(context.TODO()).Ping().Result(); err != nil {
		logger.Panicf("Redis Client Error: %v!", err)
		fmt.Println("Redis Client Error: ", err)
		return nil
	}

	return client
}
