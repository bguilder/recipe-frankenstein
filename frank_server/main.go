package main

import (
	"fmt"
	"frank_server/models"
	"frank_server/runner"
	"frank_server/scraper/allrecipes"
	"frank_server/utils"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const recipeName = "chicken and dumplings"
const numberOfRecipes = 5

func main() {
	fmt.Printf("Searching AllRecipes for: %s\n\n", recipeName)
	searchRunner := runner.SearchRunner{RecipeName: utils.UrlFormat(recipeName)}
	SearchScraper := allrecipes.SearchScraper{}

	searchRunner.Run(SearchScraper)

	recipes := []*models.Recipe{}

	for i := 0; i < numberOfRecipes; i++ {
		recipe := models.Recipe{}
		recipeRunner := runner.RecipeRunner{Recipe: recipe, RecipeLink: searchRunner.RecipeLinks[i]}
		RecipeScraper := allrecipes.RecipeScraper{}
		recipeRunner.Run(RecipeScraper)
		recipes = append(recipes, &recipeRunner.Recipe)
		time.Sleep(2 * time.Second)
	}

	allIngredients := []string{}
	for i := 0; i < len(recipes); i++ {
		for x := 0; x < len(recipes[i].Ingredients); x++ {
			allIngredients = append(allIngredients, recipes[i].Ingredients[x])
		}
	}
	calculateWordFrequency(allIngredients)
}

func hasOmittedWord(word string) bool {
	omittedWords := []string{"cup", "cups", "ounces", "ounce", "tablespoon", "tablespoons"}
	for i := 0; i < len(omittedWords); i++ {
		if strings.Contains(word, omittedWords[i]) {
			return true
		}
		if !unicode.IsLetter([]rune(word)[0]) {
			return true
		}
	}
	return false
}

func calculateWordFrequency(allIngredients []string) {

	allWordsDict := map[string]int{}
	allIngredientsArr := []string{}
	// split string
	for i := 0; i < len(allIngredients); i++ {
		splitString := strings.Split(allIngredients[i], " ")

		// remove numbers from the slice
		for x := 0; x < len(splitString); x++ {
			if _, err := strconv.Atoi(splitString[x]); err == nil {
				splitString = append(splitString[:x], splitString[x+1:]...)
			}
		}

		// remove certain words from slice
		for x := 0; x < len(splitString); x++ {
			if hasOmittedWord(splitString[x]) {
				splitString = append(splitString[:x], splitString[x+1:]...)
				x--
			}
		}

		combineString := strings.Join(splitString, " ")
		allIngredientsArr = append(allIngredientsArr, combineString)

		for x := 0; x < len(splitString); x++ {
			if _, ok := allWordsDict[splitString[x]]; ok {
				allWordsDict[splitString[x]]++
			} else {
				allWordsDict[splitString[x]] = 1
			}
		}
	}

	fmt.Printf("\n\n\n=================================Total ingredients list!=================================\n\n\n")

	for i := 0; i < len(allIngredientsArr); i++ {
		fmt.Printf("- %v\n", allIngredientsArr[i])
	}
	fmt.Printf("\n\n\n=================================Ingredients By Frequency!=================================\n\n\n")

	pairList := rankByWordCount(allWordsDict)
	for i := 0; i < len(pairList); i++ {
		fmt.Printf("ingredient: %v - %v\n", pairList[i].Key, pairList[i].Value)
	}
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
