package lists

import "errors"

var (
	ErrEmptyList     = errors.New("Empty List")
	ErrAtLeastTwo    = errors.New("List Needs At Least Two Elements")
	ErrNegativeIndex = errors.New("List Index Cannot Be Negative")
)

func MyLast(list []int) (int, error) {
	length := len(list)

	if length < 1 {
		return 0, ErrEmptyList
	}

	return list[len(list)-1], nil
}

func MyButLast(list []int) (int, error) {
	length := len(list)

	if length <= 2 {
		return 0, ErrAtLeastTwo
	}

	return list[len(list)-2], nil
}

func ElementAt(list []int, index int) (int, error) {
	length := len(list)

	if length < 1 {
		return 0, ErrEmptyList
	}

	if index < 0 {
		return 0, ErrNegativeIndex
	}

	return list[index], nil
}

func MyLength(list []int) int {
	return len(list)
}
