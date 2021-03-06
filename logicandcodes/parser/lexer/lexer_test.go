package lexer

import (
	"reflect"
	"testing"
)

type testCase struct {
	text  string
	words []string
}

var testCases = []testCase{
	{
		text:  "",
		words: []string{},
	},
	{
		text:  "A",
		words: []string{"a"},
	},
	{
		text:  "AND",
		words: []string{"and"},
	},
	{
		text:  "and a (or a b)",
		words: []string{"and", "a", "(", "or", "a", "b", ")"},
	},
	{
		text:  "AND A (OR A B)",
		words: []string{"and", "a", "(", "or", "a", "b", ")"},
	},
	{
		text:  "and a(or a b)",
		words: []string{"and", "a", "(", "or", "a", "b", ")"},
	},
	{
		text:  "(and a (or a b))",
		words: []string{"(", "and", "a", "(", "or", "a", "b", ")", ")"},
	},
}

func TestLexer(t *testing.T) {
	for _, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			actualWords := Lexer(testCase.text)
			if !reflect.DeepEqual(testCase.words, actualWords) {
				t.Fatalf("\nExpected: %v \nActual: %v", testCase.words, actualWords)
			}
		})
	}
}
