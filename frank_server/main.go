package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"frank_server/postprocessor"
	"io/ioutil"
	"log"
	"os"
)

const recipeName = "empanada"
const numberOfRecipes = 8

func main() {
	ingredients := openIngredients("ingredients_fixtures/empanada.json")
	sanitizer := postprocessor.NewSanitizer()
	p := postprocessor.NewPostProcessor(sanitizer)
	p.Run(ingredients)
}

func openIngredients(file string) []string {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Result now: %+v", result)
	return result
}

func writeToFile(ingredients []string) {
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
