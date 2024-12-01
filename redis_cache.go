package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}

func (c *redisCache) Get(ctx context.Context, key string, value interface{}) error {
	values, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return Nil
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(values), &value)
}

func (c *redisCache) Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, string(bytes), expiration).Err()
}

func (c *redisCache) Delete(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

func (c *redisCache) Keys(ctx context.Context, prefixKey string) ([]string, error) {
	res := make([]string, 0)
	iter := c.client.Scan(ctx, 0, prefixKey, 0).Iterator()
	for iter.Next(ctx) {
		res = append(res, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return res, err
	}
	return res, nil
}

func (c *redisCache) Forever(ctx context.Context, key string, value interface{}) error {
	return c.Put(ctx, key, value, time.Duration(0))
}

func (c *redisCache) Exist(ctx context.Context, key string) (int64, error) {
	return c.client.Exists(ctx, key).Result()
}

func (c *redisCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.IncrBy(ctx, key, value).Result()
}

func (c *redisCache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.DecrBy(ctx, key, value).Result()
}
	
func (c *redisCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

func (c *redisCache) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.client.SetNX(ctx, key, value, expiration).Result()
}

func (c *redisCache) Close() error {
	return c.client.Close()
}
