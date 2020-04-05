package postprocessor_test

import (
	"frank_server/postprocessor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizer_SanatizeIngredient(t *testing.T) {
	s := postprocessor.NewSanitizer()
	res := s.Sanitize("an optional 2 cherries")
	assert.Equal(t, "cherries", res)
}
