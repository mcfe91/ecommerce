package main

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(context.Context, int) ([]byte, error)
	Set(context.Context, int, any, time.Duration) error
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

func (r *RedisCache) Get(ctx context.Context, id int) ([]byte, error) {
	idStr := strconv.Itoa(id)
	val, err := r.cache.Get(ctx, idStr).Bytes()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *RedisCache) Set(ctx context.Context, id int, v any, duration time.Duration) error {
	idStr := strconv.Itoa(id)
	if err := r.cache.Set(ctx, idStr, v, duration).Err(); err != nil {
		return err
	}

	return nil
}