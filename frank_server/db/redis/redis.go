package redis

import (
	"encoding/json"
	"fmt"
	"frank_server/db"
	"frank_server/models"

	"github.com/go-redis/redis/v7"
)

type redisCache struct {
	client *redis.Client
}

// NewCache returns a new redis cache
func NewCache() db.Store {
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

func (r *redisCache) PutRecipes(searchKey string, recipes []*models.Recipe) error {
	recipeBytes, _ := json.Marshal(recipes)
	res := r.client.Set(searchKey, recipeBytes, -1)
	return res.Err()
}

func (r *redisCache) GetRecipes(searchKey string) ([]*models.Recipe, error) {
	res := r.client.Get(searchKey)
	if err := res.Err(); err != nil {
		return nil, err
	}
	resBytes, err := res.Bytes()
	if err != nil {
		return nil, err
	}
	recipes := []*models.Recipe{}
	if err := json.Unmarshal(resBytes, &recipes); err != nil {
		return nil, err
	}

	return recipes, nil
}
