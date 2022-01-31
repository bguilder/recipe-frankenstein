package main

import (
	"bufio"
	"fmt"
	"frank_server/models"
	"frank_server/runner/allrecipes"
	"log"
	"os"
)

const recipeName = "chicken marsala"
const numberOfRecipes = 2

func main() {
	scraper := allrecipes.NewAllRecipesScraper()
	recipes := scraper.GetRecipes(recipeName, numberOfRecipes)
	for _, recipe := range recipes {
		WriteToFile(recipe)
	}
}

func WriteToFile(recipe *models.Recipe) {
	fileName := fmt.Sprintf("sample/%s.txt", recipe.Title)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	datawriter.WriteString(recipe.Title + "\n\n")
	datawriter.WriteString("Ingredients:" + "\n")

	for _, data := range recipe.Ingredients {
		_, _ = datawriter.WriteString("- " + data + "\n")
	}

	datawriter.WriteString("\nDirections:" + "\n")

	for _, data := range recipe.Directions {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.WriteString("\n\n")

	datawriter.Flush()
	file.Close()
}
