package utils

import (
	"os"

	"github.com/go-redis/redis"
)

func ConnectToRedis() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"), // no password set
		DB:       0,                       // use default DB

	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil

}
