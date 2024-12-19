package lib

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetFromRedis(uri string) *redis.StringCmd {
	get := Redis().Get(context.Background(), uri)

	return get
}

func SetToRedis(uri string, encoded []byte) *redis.StatusCmd {
	set := Redis().Set(context.Background(), uri, string(encoded), 0)

	return set
}
