package listscontinued

import (
	"errors"
	"math"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/lists"
)

var (
	// ErrNotPositiveNumber - Error for factor cannot be 0 or negative
	ErrNotPositiveNumber = errors.New("factor cannot be 0 or negative")

	// ErrIndexOutOfRange - Error for index out of range
	ErrIndexOutOfRange = errors.New("index out of range")
)

// EncodeModified Problem 11
// Modified run-length encoding.
// Can't get it to omit struct value if not present
func EncodeModified(list []string) ([]lists.RLE, error) {
	packed, err := lists.Pack(list)
	if err != nil {
		return []lists.RLE{}, lists.ErrEmptyList
	}

	encoded := make([]lists.RLE, 0)

	for _, element := range packed {
		if len(element) > 1 {
			encoded = append(encoded, lists.RLE{Count: len(element), Value: string(element[0])})
		} else {
			encoded = append(encoded, lists.RLE{Value: string(element[0])})
		}
	}

	return encoded, nil
}

// DecodeModified Problem 12
// Decode a run-length encoded list.
func DecodeModified(rleList []lists.RLE) ([]string, error) {
	if len(rleList) < 1 {
		return []string{}, lists.ErrEmptyList
	}

	decoded := make([]string, 0)

	for _, element := range rleList {
		decoded = append(decoded, element.Value)
		for index := 1; index < element.Count; index++ {
			decoded = append(decoded, element.Value)
		}
	}

	return decoded, nil
}

// Problem 13 Not Solved
// Run-length encoding of a list (direct solution).

// Dupli Problem 14
// Duplicate the elements of a list.
func Dupli(list []int) ([]int, error) {
	if len(list) < 1 {
		return []int{}, lists.ErrEmptyList
	}

	duplicated := make([]int, 0)

	for _, element := range list {
		duplicated = append(duplicated, element, element)
	}

	return duplicated, nil
}

// Repli Problem 15
// Replicate the elements of a list a given number of times.
func Repli(list []int, factor int) ([]int, error) {
	if len(list) < 1 {
		return []int{}, lists.ErrEmptyList
	}

	if factor < 1 {
		return []int{}, ErrNotPositiveNumber
	}

	duplicated := make([]int, 0)

	for _, element := range list {
		for index := 0; index < factor; index++ {
			duplicated = append(duplicated, element)
		}
	}

	return duplicated, nil
}

// DropEvery Problem 16
// Drop every N'th element from a list.
func DropEvery(list []int, factor int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, lists.ErrEmptyList
	}

	if factor < 1 {
		return []int{}, ErrNotPositiveNumber
	}

	dropped := make([]int, 0)

	for index := 0; index < length; index++ {
		if (index+1)%factor != 0 {
			dropped = append(dropped, list[index])
		}
	}

	return dropped, nil
}

// Split Problem 17
// Split a list into two parts;
// the length of the first part is given.
func Split(list []int, splitPoint int) ([][]int, error) {
	length := len(list)

	if length < 1 {
		return [][]int{}, lists.ErrEmptyList
	}

	if splitPoint <= 0 || splitPoint > length {
		return [][]int{}, ErrIndexOutOfRange
	}

	divided := make([][]int, 0)

	divided = append(divided, list[:splitPoint])
	divided = append(divided, list[splitPoint:])

	return divided, nil
}

// Slice Problem 18
// Extract a slice from a list.
func Slice(list []int, startIndex int, endIndex int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, lists.ErrEmptyList
	}

	if startIndex <= 0 || startIndex > length {
		return []int{}, ErrIndexOutOfRange
	}

	if endIndex <= 0 || endIndex > length {
		return []int{}, ErrIndexOutOfRange
	}

	extract := list[startIndex-1 : endIndex]

	return extract, nil
}

// Rotate Problem 19
// Rotate a list N places to the left.
func Rotate(list []int, rotationIndex int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, lists.ErrEmptyList
	}

	if rotationIndex == 0 {
		return list, nil
	}

	if int(math.Abs(float64(rotationIndex)))%length == 0 {
		return list, nil
	}

	if rotationIndex < 0 {
		rotationIndex = (length + rotationIndex%length)
	}

	divided, err := Split(list, rotationIndex)
	if err != nil {
		return []int{}, err
	}

	rotated := make([]int, 0)
	rotated = append(rotated, divided[1]...)
	rotated = append(rotated, divided[0]...)

	return rotated, nil
}

// RemoveAt Problem 20
// Remove the K'th element from a list.
func RemoveAt(list []int, removalPoint int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, lists.ErrEmptyList
	}

	removalPoint--

	if removalPoint < 0 || removalPoint > length {
		return []int{}, ErrIndexOutOfRange
	}

	list = append(list[:removalPoint], list[removalPoint+1:]...)

	return list, nil
}
