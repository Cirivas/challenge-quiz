package redis_db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/redis/go-redis/v9"
)

type redisStore[T any] struct {
	client         *redis.Client
	collectionName string
}

func NewRedisCollection[T any](collectionName string, client *redis.Client) database.Datastore[T] {
	return &redisStore[T]{client, collectionName}
}

func (r *redisStore[T]) GetById(id string) (*T, error) {
	key := fmt.Sprintf("%s:%s", r.collectionName, id)
	value, err := r.client.JSONGet(context.Background(), key).Result()

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
