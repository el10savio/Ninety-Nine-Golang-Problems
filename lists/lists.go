package lists

import "errors"

var (
	ErrEmptyList = errors.New("Empty List")
)

func MyLast(list []int) (int, error) {
	length := len(list)

	if length < 1 {
		return 0, ErrEmptyList
	}

	return list[len(list)-1], nil
}
