package helper

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// StoreInRedis Function to store trimmedhash and value in redis
func StoreInRedis(trimmedHash string, longURL string) {
	err := client.Set(ctx, trimmedHash, longURL, 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetFromRedis Function to get value from redis
func GetFromRedis(trimmedHash string) (string, string) {
	val, err := client.Get(ctx, trimmedHash).Result()
	if err != nil {
		return "", err.Error()
	}
	return val, ""

}
