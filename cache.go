package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, any, time.Duration) error
}

type RedisCache struct {
	cache *redis.Client
}

func NewRedisCache() (*RedisCache) {
	cache := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	return &RedisCache{
		cache: cache,
	}
}

func (r *RedisCache) Get(ctx context.Context, id string) ([]byte, error) {
	val, err := r.cache.Get(ctx, id).Bytes()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, id string, v any, duration time.Duration) error {
	if err := r.cache.Set(ctx, id, v, duration).Err(); err != nil {
		return err
	}

	return nil
}