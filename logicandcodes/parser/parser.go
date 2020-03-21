package parser

import (
	"log"

	lexer "github.com/el10savio/Ninety-Nine-Golang-Problems/logicandcodes/parser/lexer"
)

type ParseTree struct {
	Type     string
	Element  string
	Children []ParseTree
}

var (
	operatorList = []string{
		"and",
		"or",
	}
)

func Parser(text string) ParseTree {
	if len(text) <= 0 {
		return ParseTree{}
	}

	var tree ParseTree
	words := lexer.Lexer(text)
	log.Println(words)

	for _, word := range words {
		tree.Type = "variable"

		if isOperator(word) {
			tree.Type = "operator"
		}

		tree.Element = word
		tree.Children = []ParseTree{}
	}

	return tree
}

func isOperator(value string) bool {
	for _, operator := range operatorList {
		if value == operator {
			return true
		}
	}
	return false
}
