package redis_demo

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetRedisClient(address, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	ctx := context.Background()

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	return client
}
