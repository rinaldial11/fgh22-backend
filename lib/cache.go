package lib

import "github.com/redis/go-redis/v9"

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.211.131:6379",
		Password: "",
		DB:       0,
	})
	return rdb
}
