package infrastracture

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return Redis{
		client: client,
	}
}

func (r *Redis) Get(key string) (string, error) {
	strcmd := r.client.Get(key)
	if err := strcmd.Err(); err != nil {
		return "", err
	}
	return strcmd.Val(), nil
}

func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	strcmd := r.client.Set(key, value, expiration)
	return strcmd.Err()
}
