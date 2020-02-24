package parser

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	sentence := "and a (or a b)"

	expectedWords := []string{"and", "a", "(", "or", "a", "b", ")"}

	actualWords := Words(sentence)

	if !reflect.DeepEqual(expectedWords, actualWords) {
		t.Fatalf("Expected: %v\n Actual: %v\n", expectedWords, actualWords)
	}
}
