package redis_test

import (
	"frank_server/cache/redis"
	"frank_server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	c := redis.NewCache()

	err := c.Save(&models.Recipe{})
	assert.NoError(t, err)
	res, err := c.Get("test")
	assert.NoError(t, err)
	assert.Equal(t, "test", res.Title)
}
