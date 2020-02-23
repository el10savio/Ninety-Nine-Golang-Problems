package parser

import (
	"log"
	"testing"
)

func TestWords(t *testing.T) {
	sentence := "and a (or a b)"
	log.Println(Words(sentence))
}
