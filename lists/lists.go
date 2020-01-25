package lists

import (
	"errors"
	"reflect"
)

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

func MyReverse(list []int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, ErrEmptyList
	}

	reversedList := make([]int, length)
	copy(reversedList, list)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		reversedList[i], reversedList[j] = reversedList[j], reversedList[i]
	}

	return reversedList, nil
}

func IsPalindrome(list []int) (bool, error) {
	reversedList, err := MyReverse(list)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(list, reversedList), nil
}

func Flatten(compositeList []interface{}) []int {
	var list []int

	for _, element := range compositeList {
		switch i := element.(type) {
		case int:
			list = append(list, i)
		case []interface{}:
			list = append(list, Flatten(i)...)
		}
	}

	return list
}
