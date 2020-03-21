package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	text  string
	words []string
}

var TestCases = []TestCase{
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
	for _, TestCase := range TestCases {
		actualWords := Lexer(TestCase.text)
		assert.Equal(t, actualWords, TestCase.words)
	}
}
