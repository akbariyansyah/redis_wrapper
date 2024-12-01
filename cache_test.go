package cache_test

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
	"github.com/akbariyansyah/redis_wrapper"
	
)

func TestNewRedisSentinel(t *testing.T) {
	miniredis, err := miniredis.Run()
	assert.NoError(t, err)

	client := cache.NewRedisSentinel([]string{miniredis.Addr()}, "")
	assert.NotNil(t, client)
}