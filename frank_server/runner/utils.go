package runner

import "strings"

// TODO: This is allreceipes specific...
const mainSelectorRecipe = "li, h1, span"
const mainSelectorSearch = "h3"
const searchPath = "/search/results/?wt="
const domain = "https://www.allrecipes.com"
const recipePath = "/recipe/258125/chicken-marsala-meatballs/"

func UrlFormat(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func normalizeString(s string) string {
	return removeExtraSpacesAndNewlines(removeUnneededWords(s))
}

func removeExtraSpacesAndNewlines(s string) string {
	return strings.Replace(strings.Join(strings.Fields(s), " "), "\n", "", -1)
}

func removeUnneededWords(s string) string {
	temp := strings.Replace(s, "Advertisement", "", -1)
	return strings.Replace(temp, "Add all ingredients to list", "", -1)
}
