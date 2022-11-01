package helpers

import (
	"github.com/go-redis/redis/v8"
)

var redisHost = "localhost:6379"
var redisPassword = ""

type RedisService interface {
}
type redisService struct {
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       0,
	})
	return client
}
