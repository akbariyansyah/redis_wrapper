package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Address  []string
	Password string
}

var Nil = errors.New("cache: value doesn't exists")

type Cache interface {
	Get(ctx context.Context, key string, value interface{}) error
	Put(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Keys(ctx context.Context, prefixKey string) ([]string, error)
	Forever(ctx context.Context, key string, value interface{}) error
	Exist(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Expire(ctx context.Context, key string, expiration time.Duration) error
	Close() error
}

func NewRedisStandalone(address []string, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     address[0],
		Password: password,
	})
}

func NewRedisSentinel(address []string, password string) *redis.Client {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: address,
		Password:      password,
	})
}

func NewRedisCluster(address []string, password string) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    address,
		Password: password,
	})
}
