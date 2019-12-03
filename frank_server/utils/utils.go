package utils

import "strings"

func UrlFormat(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}

func NormalizeString(s string) string {
	return removeExtraSpacesAndNewlines(removeUnneededWords(s))
}

func removeExtraSpacesAndNewlines(s string) string {
	return strings.Replace(strings.Join(strings.Fields(s), " "), "\n", "", -1)
}

func removeUnneededWords(s string) string {
	temp := strings.Replace(s, "Advertisement", "", -1)
	return strings.Replace(temp, "Add all ingredients to list", "", -1)
}
