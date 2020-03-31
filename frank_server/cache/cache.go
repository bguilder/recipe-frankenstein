package cache

import "frank_server/scraper"

// Simple cache for caching the results of the runner
type Store interface {
	PutRecipes(searchKey string, recipes []*scraper.Recipe) error
	GetRecipes(searchKey string) ([]*scraper.Recipe, error)
}
