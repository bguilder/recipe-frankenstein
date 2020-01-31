package redis

import (
	"fmt"
	"frank_server/cache"
	"frank_server/models"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	client *redis.Client
}

// NewCache returns a new redis cache
func NewCache() cache.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return &redisCache{
		client: client,
	}
}

func (r *redisCache) Save(recipe *models.Recipe) error {
	res := r.client.Set("test", "foo", -1)
	return res.Err()
}

func (r *redisCache) Get(recipeName string) (*models.Recipe, error) {
	res := r.client.Get(recipeName)
	if err := res.Err(); err != nil {
		return nil, err
	}
	return &models.Recipe{Title: recipeName}, nil
}
