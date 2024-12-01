package cache

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClusterCache struct {
	client *redis.ClusterClient
}

func NewRedisClusterCache(client *redis.ClusterClient) Cache {
	return &redisClusterCache{client: client}
}

func (c *redisClusterCache) Get(ctx context.Context, key string, value interface{}) error {
	values, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return Nil
	} else if err != nil {
		return err
	}
	return json.Unmarshal([]byte(values), &value)
}

func (c *redisClusterCache) Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, string(bytes), expiration).Err()
}

func (c *redisClusterCache) Delete(ctx context.Context, keys ...string) error {
	errMessages := make([]string, 0)

	for _, key := range keys {
		if err := c.client.Del(ctx, key).Err(); err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}

	if len(errMessages) > 0 {
		return errors.New(strings.Join(errMessages, ", "))
	}
	return nil
}

func (c *redisClusterCache) Keys(ctx context.Context, prefixKey string) ([]string, error) {
	res := make([]string, 0)
	err := c.client.ForEachMaster(ctx, func(ctx context.Context, client *redis.Client) error {
		iter := c.client.Scan(ctx, 0, prefixKey, 0).Iterator()
		for iter.Next(ctx) {
			res = append(res, iter.Val())
		}

		if err := iter.Err(); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *redisClusterCache) Forever(ctx context.Context, key string, value interface{}) error {
	return c.Put(ctx, key, value, time.Duration(0))
}

func (c *redisClusterCache) Exist(ctx context.Context, key string) (int64, error) {
	return c.client.Exists(ctx, key).Result()
}

func (c *redisClusterCache) IncrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.IncrBy(ctx, key, value).Result()
}

func (c *redisClusterCache) DecrBy(ctx context.Context, key string, value int64) (int64, error) {
	return c.client.DecrBy(ctx, key, value).Result()
}

func (c *redisClusterCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}
func (c *redisClusterCache) Close() error {
	return c.client.Close()
}
