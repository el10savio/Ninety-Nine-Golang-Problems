package lexer

import (
	"strings"
)

func Lexer(text string) []string {
	if len(text) > 0 {
		text = strings.ToLower(text)

		text = strings.ReplaceAll(text, "(", " ( ")
		text = strings.ReplaceAll(text, ")", " ) ")

		text = strings.ReplaceAll(text, "  ", " ")
		text = strings.Trim(text, " ")

		words := strings.Split(text, " ")
		return words
	}

	return []string{}
}
