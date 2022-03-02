// Package cache
// @Author:        asus
// @Description:   $
// @File:          redis
// @Data:          2022/3/218:43
//
package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"new-project/pkg/config"
)

func InitRedis(option config.Reids) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     option.Host,
		Password: option.Password,
		DB:       option.DefaultDB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
