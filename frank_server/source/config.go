package source

// RecipeSourceConfig is config specific to a recipe source that is used by the scraper
type RecipeSourceConfig struct {
	MainSelector string
}

// LinkSourceConfig is config specific to a link source that is used by the scraper
type LinkSourceConfig struct {
	MainSelector string
	LinkPath     string
	Domain       string
}
