package models

import "github.com/go-redis/redis"

var (
	RedisClient *redis.Client
)

type(
	Pair struct {
		Key string
		Value string
	}
)
