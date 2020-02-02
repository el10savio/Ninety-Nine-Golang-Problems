package listsagain

import (
	"math/rand"
	"time"

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

func RndSelect(list []int, count int) ([]int, error) {
	length := len(list)

	if length < 1 {
		return []int{}, definitions.ErrEmptyList
	}

	if count < 0 {
		return []int{}, definitions.ErrNegativeNumber
	}

	if count == 0 {
		return []int{}, nil
	}

	items := make([]int, count)
	rand.Seed(time.Now().UnixNano())

	for index := 0; index < count; index++ {
		element := list[rand.Intn(length)]
		items[index] = element
	}

	return items, nil
}

func DiffSelect(count int, max int) ([]int, error) {
	if count < 0 {
		return []int{}, definitions.ErrNegativeNumber
	}

	if count == 0 {
		return []int{}, nil
	}

	if max == 0 {
		return []int{}, definitions.ErrZero
	}

	items := make([]int, count)
	rand.Seed(time.Now().UnixNano())

	if max > 0 {
		for index := 0; index < count; index++ {
			element := rand.Intn(max) + 1
			items[index] = element
		}
	} else {
		for index := 0; index < count; index++ {
			element := -1 * (rand.Intn(-1*max) + 1)
			items[index] = element
		}
	}
	return items, nil
}
