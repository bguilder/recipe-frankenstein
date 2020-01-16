package postprocessor

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type PostProcessor struct {
	sanitizer Sanitizer
}

func NewPostProcessor(sanitizer Sanitizer) PostProcessor {
	return PostProcessor{sanitizer: NewSanitizer()}
}

func (p *PostProcessor) Run(ingredients []string) PairList {
	return p.calculateWordFrequency(ingredients)
}

func (p *PostProcessor) calculateWordFrequency(allIngredients []string) PairList {

	allWordsDict := map[string]int{}
	allIngredientsArr := []string{}
	allIngredientsDict := map[string]int{}
	// split string
	for i := 0; i < len(allIngredients); i++ {
		initialSplit := strings.Split(allIngredients[i], " ")

		splitString := []string{}

		for _, word := range initialSplit {
			ingredient := p.sanitizer.RemovePunctuation(word)
			ingredient = p.sanitizer.ToLower(ingredient)
			splitString = append(splitString, ingredient)
		}

		// remove numbers from the slice
		for x := 0; x < len(splitString); x++ {
			if _, err := strconv.Atoi(splitString[x]); err == nil {
				splitString = append(splitString[:x], splitString[x+1:]...)
			}
		}

		// remove certain words from slice
		for x := 0; x < len(splitString); x++ {
			if p.sanitizer.hasStopWord(splitString[x]) {
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

		if _, ok := allIngredientsDict[combineString]; ok {
			allIngredientsDict[combineString]++
		} else {
			allIngredientsDict[combineString] = 1
		}

	}

	fmt.Printf("\n\n\n=================================allIngredientsArr!=================================\n\n\n")

	for i := 0; i < len(allIngredientsArr); i++ {
		fmt.Printf("- %v\n", allIngredientsArr[i])
	}

	fmt.Printf("\n\n\n=================================pairList1!=================================\n\n\n")

	pairList1 := rankByMultipleWordsCount(allIngredientsDict)
	for i := 0; i < len(pairList1); i++ {
		fmt.Printf("ingredient: %v - %v\n", pairList1[i].Key, pairList1[i].Value)
	}

	fmt.Printf("\n\n\n=================================pairList!=================================\n\n\n")

	pairList := rankByWordCount(allWordsDict)
	for i := 0; i < len(pairList); i++ {
		fmt.Printf("ingredient: %v - %v\n", pairList[i].Key, pairList[i].Value)
	}

	return pairList1
}

func rankByMultipleWordsCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
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
