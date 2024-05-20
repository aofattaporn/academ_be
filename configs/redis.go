package configs

import (
	"github.com/go-redis/redis"
)

func ConnectRedis() *redis.Client {

	url := "redis://red-couq3r821fec73brsf50:6379"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
