package postprocessor

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// Sanitizer filters the results of the scraper
type Sanitizer struct {
	StopWords map[string]interface{}
}

func loadStopWords() map[string]interface{} {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(stopWords), &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// NewSanitizer returns a Sanitizer
func NewSanitizer() Sanitizer {
	// load the list of stop words
	return Sanitizer{StopWords: loadStopWords()}
}

func (s *Sanitizer) hasStopWord(word string) bool {
	for i := 0; i < len(s.StopWords); i++ {
		if _, ok := s.StopWords[word]; ok {
			fmt.Printf("has stop word... %s", word)
			return true
		}
	}
	return false
}

// Sanitize removes stop words, removes punctuation and lower cases
func (s *Sanitizer) Sanitize(ingredient string) string {
	// remove stop words
	splitIngredient := strings.Split(ingredient, " ")
	for i := 0; i < len(splitIngredient); i++ {
		if s.hasStopWord(splitIngredient[i]) {
			splitIngredient = append(splitIngredient[:i], splitIngredient[i+1:]...)
			i--
		}
	}
	ingredient = strings.Join(splitIngredient, " ")

	// remove everything that is not a letter
	runes := []rune(ingredient)
	for i := 0; i < len(runes); i++ {
		if !unicode.IsLetter(runes[i]) {
			runes = append(runes[:i], runes[i+1:]...)
			i--
		}
	}
	// lower case
	return strings.ToLower(string(runes))
}
