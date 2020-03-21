package parser

import (
	lexer "github.com/el10savio/Ninety-Nine-Golang-Problems/logicandcodes/parser/lexer"
)

type ParseTree struct {
	Type     string
	Element  string
	Children []*ParseTree
}

var (
	operatorList    = []string{"and", "or"}
	punctuationList = []string{"(", ")"}
)

func Parser(text string) ParseTree {
	if len(text) <= 0 {
		return ParseTree{}
	}

	var tree ParseTree
	words := lexer.Lexer(text)

	for _, word := range words {
		tree.Element = word

		if !isOperator(word) {
			tree.Type = "variable"
			continue
		}

		tree.Type = "operator"
		tree.Children = []*ParseTree{}
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
