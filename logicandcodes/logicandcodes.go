package logicandcodes

import (
	"log"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
	"github.com/el10savio/Ninety-Nine-Golang-Problems/listsagain"
)

func GrayCodes(number int) ([]string, error) {
	if number < 1 {
		return []string{}, definitions.ErrNotPositiveNumber
	}

	baseValues := []string{"0", "1"}

	if number == 1 {
		return baseValues, nil
	}

	combinations, err := listsagain.Combinations(number, baseValues)
	if err != nil {
		return []string{}, err
	}

	log.Println(combinations)

	return []string{}, nil
}
