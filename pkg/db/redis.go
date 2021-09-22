package db

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	expire time.Duration
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewRedisCache(duration time.Duration) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisCache{
		client: client,
		expire: duration,
	}, nil
}

func (c *RedisCache) Set(key string, value interface{}) {
	json, err := json.Marshal(value)
	if err != nil {
		log.Fatalf("error %v", err)
	}

	c.client.Set(context.Background(), key, json, c.expire*time.Second)
}

func (c *RedisCache) Get(key string) (string, error) {
	val, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
