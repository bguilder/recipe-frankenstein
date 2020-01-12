package postprocessor_test

import (
	"frank_server/postprocessor"
	"testing"
)

func TestSanitizer(t *testing.T) {
	sanitzier := &postprocessor.Sanitizer{}
	res := sanitzier.RemovePunctuation("test")
	if res != "test" {
		t.Fail()
	}
}
