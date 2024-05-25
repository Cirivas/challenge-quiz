package redis_db

import (
	"github.com/cirivas/challenge-quiz/infrastructure/database"
	"github.com/redis/go-redis/v9"
)

type redisSt struct {
	client *redis.Client
}

func NewRedisClient() database.DatastoreClient {
	return &redisSt{}
}

func (r *redisSt) Client() interface{} {
	return r.client
}

func (r *redisSt) Connect() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r.client = rdb
	return nil
}

func (r *redisSt) Close() error {
	return r.client.Close()
}
