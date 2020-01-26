package lists

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrEmptyList - Error for empty list
	ErrEmptyList = errors.New("empty list")

	// ErrAtLeastTwo - Error for list needs at least two elements
	ErrAtLeastTwo = errors.New("list needs at least two elements")

	// ErrNegativeIndex - Error for list index cannot be negative
	ErrNegativeIndex = errors.New("list index cannot be negative")
)

// RLE is the type used in
// Encode(list []string) ([]RLE, error)
type RLE struct {
	count int
	value string
}

// MyLast Problem 1
// Find the last element of a list.
func MyLast(list []int) (int, error) {
	length := len(list)

	if length < 1 {
		return 0, ErrEmptyList
	}

	return list[len(list)-1], nil
}

// MyButLast Problem 2
// Find the last but one element of a list.
func MyButLast(list []int) (int, error) {
	length := len(list)

	if length <= 2 {
		return 0, ErrAtLeastTwo
	}

	return list[len(list)-2], nil
}

// ElementAt Problem 3
// Find the K'th element of a list.
// The first element in the list is index 1.
func ElementAt(list []int, index int) (int, error) {
	length := len(list)

	if length < 1 {
		return 0, ErrEmptyList
	}

	if index < 0 {
		return 0, ErrNegativeIndex
	}

	index--
	return list[index], nil
}

// MyLength Problem 4
// Find the number of elements of a list.
func MyLength(list []int) int {
	return len(list)
}

// MyReverse Problem 5
// Reverse a list.
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

// IsPalindrome Problem 6
// Find out whether a list is a palindrome.
// A palindrome can be read forward or backward;
func IsPalindrome(list []int) (bool, error) {
	reversedList, err := MyReverse(list)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(list, reversedList), nil
}

// Flatten Problem 7
// Flatten a nested list structure.
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

// Compress Problem 8
// Eliminate consecutive duplicates of list elements.
func Compress(list []string) ([]string, error) {
	length := len(list)

	if length < 1 {
		return []string{}, ErrEmptyList
	}

	present := make(map[string]bool)
	compressed := make([]string, 0)

	for _, element := range list {
		if !present[element] {
			present = make(map[string]bool)
			present[element] = true
			compressed = append(compressed, element)
		}
	}

	return compressed, nil
}

// Pack Problem 9
// Pack consecutive duplicates of list elements into sublists.
// If a list contains repeated elements they should be placed in separate sublists.
func Pack(list []string) ([]string, error) {
	length := len(list)

	if length < 1 {
		return []string{}, ErrEmptyList
	}

	present := make(map[string]bool)
	compressed := make([]string, 0)

	joined := list[0]
	present[joined] = true

	for index := 1; index < length; index++ {
		element := list[index]

		if !present[element] {
			compressed = append(compressed, joined)
			present = make(map[string]bool)
			joined = element
		} else {
			joined = fmt.Sprintf("%s%s", joined, element)
		}

		present[element] = true
	}
	compressed = append(compressed, joined)

	return compressed, nil
}

// Encode Problem 10
// Run-length encoding of a list.
func Encode(list []string) ([]RLE, error) {
	packed, err := Pack(list)
	if err != nil {
		return []RLE{}, ErrEmptyList
	}

	encoded := make([]RLE, 0)

	for _, element := range packed {
		encoded = append(encoded, RLE{count: len(element), value: string(element[0])})
	}

	return encoded, nil
}
