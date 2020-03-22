package parser

import (
	lexer "github.com/el10savio/Ninety-Nine-Golang-Problems/logicandcodes/parser/lexer"
)

type ParseTree struct {
	Type     string
	Element  string
	Children []ParseTree
}

var (
	operatorList    = []string{"and", "or"}
	punctuationList = []string{"(", ")"}
)

func Parser(text string) ParseTree {
	if len(text) <= 0 {
		return ParseTree{}
	}

	words := lexer.Lexer(text)
	tree := generateTree(words)

	return tree
}

func generateTree(words []string) (tree ParseTree) {
	if len(words) <= 0 {
		return
	}

	tree.Type = "operator"
	tree.Element = words[0]
	tree.Children = []ParseTree{}

	if !isOperator(words[0]) {
		tree.Type = "variable"
	}

	for _, word := range words[1:] {
		if !isOperator(word) {
			tree.Children = append(tree.Children, generateNode(word))
		}
	}

	return
}

func generateNode(word string) (node ParseTree) {
	if len(word) <= 0 {
		return
	}

	node.Element = word
	node.Type = "operator"
	node.Children = []ParseTree{}

	if !isOperator(word) {
		node.Type = "variable"
	}

	return
}

func isOperator(value string) bool {
	for _, operator := range operatorList {
		if value == operator {
			return true
		}
	}
	return false
}
