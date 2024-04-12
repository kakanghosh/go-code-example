package redis_demo

import "github.com/redis/go-redis/v9"

func GetRedisClient(address, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}
