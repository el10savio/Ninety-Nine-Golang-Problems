package listsagain

import (
	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

func InsertAt(value int, list []int, insertPoint int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, definitions.ErrEmptyList
	}

	insertPoint--

	if insertPoint <= 0 || insertPoint > length {
		return []int{}, definitions.ErrIndexOutOfRange
	}

	newList := append(list, 0)
	copy(newList[insertPoint+1:], newList[insertPoint:])
	newList[insertPoint] = value

	return newList, nil
}

func Range(startValue int, endValue int) []int {
	if startValue == endValue {
		return []int{}
	}

	list := make([]int, 0)

	if startValue < endValue {
		for value := startValue; value <= endValue; value++ {
			list = append(list, value)
		}
	} else {
		for value := startValue; value >= endValue; value-- {
			list = append(list, value)
		}

	}

	return list
}
