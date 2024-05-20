package handlers

import (
	"academ_be/configs"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	redisClient = configs.ConnectRedis()
}
