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

	tree = generateNode(words[0])
	node := make([]ParseTree, 0)

	for index, word := range words[1:] {
		if isPunctuation(word) {
			continue
		}

		if isOperator(word) {
			node = append(node, generateTree(words[index+1:]))
			break
		}

		node = append(node, generateNode(word))
	}

	tree.Children = node
	return
}

func generateNode(word string) (node ParseTree) {
	if len(word) <= 0 {
		return
	}

	if isPunctuation(word) {
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

func isPunctuation(value string) bool {
	for _, punctuation := range punctuationList {
		if value == punctuation {
			return true
		}
	}
	return false
}
