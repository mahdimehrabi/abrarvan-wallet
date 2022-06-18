package infrastracture

import (
	"bytes"
	"challange/app/models"
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
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
	ctx := context.Background()
	strcmd := r.client.Get(ctx, key)
	if err := strcmd.Err(); err != nil {
		return "", err
	}
	return strcmd.Val(), nil
}

func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	strcmd := r.client.Set(ctx, key, value, expiration)
	return strcmd.Err()
}

func (r *Redis) DecreaseConsumerCount(code string) (consumerCount int, err error) {
	ctx := context.Background()
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		bJson, err := tx.Get(ctx, "code_"+code).Bytes()
		if err != nil && err != redis.Nil {
			return err
		}
		bReader := bytes.NewReader(bJson)

		code := models.Code{}
		err = code.FromJSON(bReader)
		if err != nil {
			return err
		}
		code.ConsumerCount--
		consumerCount = code.ConsumerCount

		var buff bytes.Buffer
		err = code.ToJSON(&buff)
		if err != nil {
			return err
		}

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, code.Code, buff.String(), 0)
			return nil
		})
		return err
	}

	// Retry if the key has been changed.
	for i := 0; i < 200; i++ {
		err := r.client.Watch(ctx, txf, "code_"+code)
		if err == nil {
			// Success.
			return consumerCount, nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return 0, err
	}

	return 0, errors.New("increment reached maximum number of retries")
}
