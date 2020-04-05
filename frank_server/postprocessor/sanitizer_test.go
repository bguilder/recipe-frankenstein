package postprocessor_test

import (
	"frank_server/postprocessor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizer_SanatizeIngredient(t *testing.T) {
	s := postprocessor.NewSanitizer()
	testCases := map[string]string{
		"2 large cherries":        "cherries",
		"2 quarts popped popcorn": "popped popcorn",
		"7(*^4 😆woop woop":        "woop woop",
	}
	for testCase, expected := range testCases {
		assert.Equal(t, expected, s.Sanitize(testCase))
	}
}
