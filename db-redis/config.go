package redis

import (
	"os"

	redis "github.com/go-redis/redis"
)

// Setup initializes a redis client
func Setup() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0, // use default DB
	})

	// commands returns "PONG", tests if
	// connection alive
	_, err := client.Ping().Result()
	return client, err
}
