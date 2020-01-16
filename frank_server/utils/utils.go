package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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

// opens a static file of ingredients
// ../../ingredients_fixtures/empanada.json
func OpenIngredients(file string) []string {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func WriteToFile(ingredients []string) {
	file, err := os.OpenFile("raw_ingredients.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range ingredients {
		_, _ = datawriter.WriteString(data + "\n")
	}
	datawriter.Flush()
	file.Close()
}
