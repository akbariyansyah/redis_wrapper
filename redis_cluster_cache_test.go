package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"github.com/akbariyansyah/redis_wrapper"
)

type redisClusterCacheTestSuite struct {
	suite.Suite
	ctx        context.Context
	ttl        time.Duration
	redis      *redis.ClusterClient
	redisCache cache.Cache
}

func TestRedisClusterCacheTestSuite(t *testing.T) {
	suite.Run(t, new(redisClusterCacheTestSuite))
}

func (t *redisClusterCacheTestSuite) SetupTest() {
	t.ctx = context.Background()
	t.ttl = time.Second * 60
	miniredis, err := miniredis.Run()
	if err != nil {
		t.Errorf(err, "miniredis")
	}
	t.redis = cache.NewRedisCluster([]string{miniredis.Addr()}, "")
	t.redisCache = cache.NewRedisClusterCache(t.redis)
}

func (t *redisClusterCacheTestSuite) TearDownTest() {
	t.redis.Close()
}

func (t *redisClusterCacheTestSuite) TestGet_Nil() {
	var testResult int
	err := t.redisCache.Get(t.ctx, "key1", testResult)

	t.Equal(cache.Nil, err)
}

func (t *redisClusterCacheTestSuite) TestGet_NoError() {
	mockData := 10

	err := t.redisCache.Put(t.ctx, "key1", mockData, t.ttl)
	t.NoError(err)

	var testResult int
	err = t.redisCache.Get(t.ctx, "key1", testResult)

	t.NoError(err)
	t.NotEqual(mockData, testResult)
}

func (t *redisClusterCacheTestSuite) TestPut_Error() {
	err := t.redisCache.Put(t.ctx, "key1", make(chan int), t.ttl)
	t.Error(err)
}

func (t *redisClusterCacheTestSuite) TestDelete() {
	err := t.redisCache.Delete(t.ctx, "key1")
	t.NoError(err)
}

func (t *redisClusterCacheTestSuite) TestKeys_NotExists() {
	keys, err := t.redisCache.Keys(t.ctx, "key1")
	t.NoError(err)
	t.Len(keys, 0)
}

func (t *redisClusterCacheTestSuite) TestKeys_Exists() {
	err := t.redisCache.Put(t.ctx, "key1", "xxxx", t.ttl)
	t.NoError(err)

	res, err := t.redisCache.Keys(t.ctx, "key1")
	t.NoError(err)
	t.Equal([]string([]string{"key1"}), res)
}

func (t *redisClusterCacheTestSuite) TestForever() {
	err := t.redisCache.Forever(t.ctx, "key1", "xxxx")
	t.NoError(err)
}

func (t *redisClusterCacheTestSuite) TestExist_NotExists() {
	res, err := t.redisCache.Exist(t.ctx, "key1")
	t.NoError(err)
	t.Equal(int64(0), res)
}

func (t *redisClusterCacheTestSuite) TestIncrBy() {
	total, err := t.redisCache.IncrBy(t.ctx, "key1", int64(1000))
	t.NoError(err)
	t.Equal(int64(1000), total)
}

func (t *redisClusterCacheTestSuite) TestDecrBy() {
	total, err := t.redisCache.DecrBy(t.ctx, "key1", int64(1000))
	t.NoError(err)
	t.Equal(int64(-1000), total)
}

func (t *redisClusterCacheTestSuite) TestExpire() {
	err := t.redisCache.Expire(t.ctx, "key1", t.ttl)
	t.NoError(err)
}

func (t *redisClusterCacheTestSuite) TestClose() {
	err := t.redisCache.Close()
	t.NoError(err)
}