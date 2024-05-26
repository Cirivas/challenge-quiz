package redis_db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/redis/go-redis/v9"
)

type redisStore[T any] struct {
	client *redis.Client
}

func NewRedisCollection[T any](client *redis.Client) database.Datastore[T] {
	return &redisStore[T]{client}
}

func (r *redisStore[T]) GetById(id string) (*T, error) {
	value, err := r.client.JSONGet(context.Background(), id).Result()

	if err != nil {
		return nil, err
	}

	if value == "" {
		fmt.Printf("No value found for id %s\n", id)
		return nil, nil
	}

	var parsedToType []T
	if err = json.Unmarshal([]byte(value), &parsedToType); err != nil {
		fmt.Printf("Marshal error: %v; value: %#v\n", err, value)
		return nil, err
	}

	return &parsedToType[0], nil
}

func (r *redisStore[T]) Get(...database.SearchField) ([]T, error) {
	return nil, nil
}

func (r *redisStore[T]) Save(val T) error {
	return nil
}
