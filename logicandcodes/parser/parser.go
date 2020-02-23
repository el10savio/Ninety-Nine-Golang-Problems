package parser

import "strings"

func Words(sentence string) []string {
	sentence = strings.Trim(sentence, " ")

	sentence = strings.ReplaceAll(sentence, "(", "( ")
	sentence = strings.ReplaceAll(sentence, ")", " )")

	words := strings.Split(sentence, " ")
	return words
}
