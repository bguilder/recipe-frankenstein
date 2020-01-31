package cache

import "frank_server/models"

// Cache saves recently searched recipes for a configured amount of time
type Cache interface {
	Save(recipe *models.Recipe) error
	Get(recipeName string) (*models.Recipe, error)
}
