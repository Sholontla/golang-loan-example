package redis

import "github.com/go-redis/redis/v9"

var Cache *redis.Client

func SetUpRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}
