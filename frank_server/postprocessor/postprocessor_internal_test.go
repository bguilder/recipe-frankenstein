package postprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizer_CalculateWordFrequency(t *testing.T) {
	post := NewPostProcessor()
	words := []string{"lettuce", "tomato", "onion", "tomato", "onion", "onion"}
	res := post.calculateWordFrequency(words)
	expected := IngredientFrequencyList{
		IngredientFrequency{Name: "onion", Count: 3},
		IngredientFrequency{Name: "tomato", Count: 2},
		IngredientFrequency{Name: "lettuce", Count: 1},
	}
	assert.Equal(t, expected, res)
}
