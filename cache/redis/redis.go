package rediscache

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cache *redis.Client
}

func NewRedisCache(ctx context.Context, addr, password string, db int) (*RedisCache, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	connResp := cache.Ping(ctx)
	if connResp.Val() != "PONG" {
		return nil, errors.New("cannot connect to redis: no ping")
	}

	return &RedisCache{
		cache: cache,
	}, nil
}

func (rc *RedisCache) Insert(ctx context.Context, key, value string) error {
	err := rc.cache.Set(ctx, key, value, 0).Err()

	if err != nil {
		return fmt.Errorf("cannot write to redis: %w", err)
	}

	return nil
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	redisResponse := rc.cache.Get(ctx, key)

	if err := redisResponse.Err(); err != nil {
		return "", fmt.Errorf("cannot get value from redis: %w", err)
	}

	result, err := redisResponse.Result()
	if err != nil {
		return "", fmt.Errorf("cannot get value from redis: %w", err)
	}

	return result, nil
}

func (rc *RedisCache) Delete(ctx context.Context, key string) error {
	redisResponse := rc.cache.Del(ctx, key)

	if err := redisResponse.Err(); err != nil {
		return fmt.Errorf("cannot delete value in redis: %w", err)
	}

	return nil
}
