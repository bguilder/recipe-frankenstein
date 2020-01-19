package postprocessor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

// Sanitizer filters the results of the scraper
type Sanitizer struct {
	StopWords map[string]interface{}
}

func loadStopWords() map[string]interface{} {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	jsonFile, err := os.Open("./stop_words.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened stop_words.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
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

// TODO:
// - Remove punctuation
// - Lower case everything
// -

func (s *Sanitizer) hasStopWord(word string) bool {

	for i := 0; i < len(s.StopWords); i++ {
		// Make sure the first characters is a letter
		if word != "" {
			if !unicode.IsLetter([]rune(word)[0]) {
				return true
			}
		}
		if _, ok := s.StopWords[word]; ok {
			fmt.Printf("has stop word... %s", word)
			return true
		}
	}
	return false
}

func (s *Sanitizer) RemovePunctuation(word string) string {
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		if unicode.IsPunct(runes[i]) {
			runes = append(runes[:i], runes[i+1:]...)
			i--
		}
	}
	return string(runes)
}

func (s *Sanitizer) ToLower(word string) string {
	return strings.ToLower(word)
}
