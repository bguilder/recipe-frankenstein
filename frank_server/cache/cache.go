package cache

import "frank_server/models"

// Simple cache for caching the results of the runner
type Store interface {
	PutRecipes(searchKey string, recipes []*models.Recipe) error
	GetRecipes(searchKey string) ([]*models.Recipe, error)
}
