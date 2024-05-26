package redis_db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/redis/go-redis/v9"
)

type Keyer interface {
	Key() string
}

type redisStore[T Keyer] struct {
	client         *redis.Client
	collectionName string
}

func NewRedisCollection[T Keyer](collectionName string, client *redis.Client) database.Datastore[T] {
	return &redisStore[T]{client, collectionName}
}

func (r *redisStore[T]) GetById(id string) (*T, error) {
	log.Println("Called GetById with value", id)
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
	// pretty inefficient: no support for RediSearch on go-redis :(
	matcher := fmt.Sprintf("%s:*", r.collectionName)
	scanResult, _, _ := r.client.Scan(context.Background(), 0, matcher, 0).Result()

	results := make([]T, len(scanResult))
	for i, v := range scanResult {
		// remove collection name
		sliced := strings.Split(v, ":")
		value, _ := r.GetById(strings.Join(sliced[1:], ":"))
		results[i] = *value
	}

	return results, nil
}

func (r *redisStore[T]) Save(val T) error {
	key := fmt.Sprintf("%s:%s", r.collectionName, val.Key())

	log.Println("Saving in path", key, val)

	if _, err := r.client.JSONSet(context.Background(), key, "$", val).Result(); err != nil {
		return err
	}

	return nil
}
