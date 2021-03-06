package redisClient

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
)

type Redis struct {
	Config *redis.Options
	Client *redis.Client
}

func RedisAdapter(config *redis.Options) *Redis {
	conn := NewClient()
	return &Redis{Config: config, Client: conn}
}

func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rdb.AddHook(apmgoredis.NewHook())
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	return rdb
}
