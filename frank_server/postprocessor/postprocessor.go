package postprocessor

import (
	"sort"
)

// IngredientFrequency tracks the count of ingredients
type IngredientFrequency struct {
	Name  string
	Count int
}

// IngredientFrequencyList implements the sort interface
type IngredientFrequencyList []IngredientFrequency

func (x IngredientFrequencyList) Len() int           { return len(x) }
func (x IngredientFrequencyList) Less(i, j int) bool { return x[i].Count < x[j].Count }
func (x IngredientFrequencyList) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// PostProcessor sanitizes and sorts the ingredients
type PostProcessor struct {
	sanitizer Sanitizer
}

// NewPostProcessor returns a new PostProcesor
func NewPostProcessor() PostProcessor {
	return PostProcessor{sanitizer: NewSanitizer()}
}

// Run runs the processor on the array of ingredients
func (p *PostProcessor) Run(ingredients []string) IngredientFrequencyList {
	return p.calculateWordFrequency(ingredients)
}

func (p *PostProcessor) calculateWordFrequency(allIngredients []string) IngredientFrequencyList {
	allIngredientsDict := map[string]int{}
	for _, ingredient := range allIngredients {
		sanitizedIngredient := p.sanitizer.Sanitize(ingredient)
		if _, ok := allIngredientsDict[sanitizedIngredient]; ok {
			allIngredientsDict[sanitizedIngredient]++
		} else {
			allIngredientsDict[sanitizedIngredient] = 1
		}
	}
	return rankByWordCount(allIngredientsDict)
}

func rankByWordCount(ingredients map[string]int) IngredientFrequencyList {
	result := make(IngredientFrequencyList, len(ingredients))
	i := 0
	for k, v := range ingredients {
		result[i] = IngredientFrequency{k, v}
		i++
	}
	sort.Sort(sort.Reverse(result))
	return result
}
