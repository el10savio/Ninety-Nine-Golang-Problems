package parser

import (
	"strings"
)

type ParseTree struct {
	operator string
	left     *ParseTree
	right    *ParseTree
}

func Words(sentence string) []string {
	sentence = strings.ReplaceAll(sentence, "(", " ( ")
	sentence = strings.ReplaceAll(sentence, ")", " ) ")

	sentence = strings.ReplaceAll(sentence, "  ", " ")
	sentence = strings.Trim(sentence, " ")

	words := strings.Split(sentence, " ")

	return words
}
