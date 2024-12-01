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

type redisCacheTestSuite struct {
	suite.Suite
	ctx        context.Context
	ttl        time.Duration
	miniredis  *miniredis.Miniredis
	redis      *redis.Client
	redisCache cache.Cache
}

func TestRedisCacheTestSuite(t *testing.T) {
	suite.Run(t, new(redisCacheTestSuite))
}

func (t *redisCacheTestSuite) SetupTest() {
	var err error
	t.ctx = context.Background()
	t.ttl = time.Second * 60
	t.miniredis, err = miniredis.Run()
	if err != nil {
		t.Errorf(err, "miniredis")
	}
	t.redis = cache.NewRedisStandalone([]string{t.miniredis.Addr()}, "")
	t.redisCache = cache.NewRedisCache(t.redis)
}

func (t *redisCacheTestSuite) TearDownTest() {
	t.redis.Close()
}

func (t *redisCacheTestSuite) TestGet_Nil() {
	var testResult int
	err := t.redisCache.Get(t.ctx, "key1", testResult)

	t.Equal(cache.Nil, err)
}

func (t *redisCacheTestSuite) TestGet_NoError() {
	mockData := 10

	err := t.redisCache.Put(t.ctx, "key1", mockData, t.ttl)
	t.NoError(err)

	var testResult int
	err = t.redisCache.Get(t.ctx, "key1", testResult)

	t.NoError(err)
	t.NotEqual(mockData, testResult)
}

func (t *redisCacheTestSuite) TestPut_Error() {
	err := t.redisCache.Put(t.ctx, "key1", make(chan int), t.ttl)
	t.Error(err)
}

func (t *redisCacheTestSuite) TestDelete() {
	err := t.redisCache.Delete(t.ctx, "key1")
	t.NoError(err)
}

func (t *redisCacheTestSuite) TestKeys_NotExists() {
	keys, err := t.redisCache.Keys(t.ctx, "key1")
	t.NoError(err)
	t.Len(keys, 0)
}

func (t *redisCacheTestSuite) TestKeys_Exists() {
	err := t.redisCache.Put(t.ctx, "key1", "xxxx", t.ttl)
	t.NoError(err)

	res, err := t.redisCache.Keys(t.ctx, "key1")
	t.NoError(err)
	t.Equal([]string([]string{"key1"}), res)
}

func (t *redisCacheTestSuite) TestForever() {
	err := t.redisCache.Forever(t.ctx, "key1", "xxxx")
	t.NoError(err)
}

func (t *redisCacheTestSuite) TestExist_NotExists() {
	res, err := t.redisCache.Exist(t.ctx, "key1")
	t.NoError(err)
	t.Equal(int64(0), res)
}

func (t *redisCacheTestSuite) TestIncrBy() {
	total, err := t.redisCache.IncrBy(t.ctx, "key1", int64(1000))
	t.NoError(err)
	t.Equal(int64(1000), total)
}

func (t *redisCacheTestSuite) TestDecrBy() {
	total, err := t.redisCache.DecrBy(t.ctx, "key1", int64(1000))
	t.NoError(err)
	t.Equal(int64(-1000), total)
}

func (t *redisCacheTestSuite) TestExpire() {
	err := t.redisCache.Expire(t.ctx, "key1", t.ttl)
	t.NoError(err)
}

func (t *redisCacheTestSuite) TestClose() {
	err := t.redisCache.Close()
	t.NoError(err)
}
