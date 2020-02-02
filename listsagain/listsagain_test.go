package listsagain

import (
	"reflect"
	"testing"

	"github.com/el10savio/Ninety-Nine-Golang-Problems/definitions"
)

func TestInsertAt(t *testing.T) {
	list := []int{1, 2, 3, 4}
	insertPoint := 2
	value := 7

	expectedInsertAt := []int{1, 7, 2, 3, 4}
	var expectedErr error

	actualInsertAt, actualErr := InsertAt(value, list, insertPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedInsertAt, actualInsertAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedInsertAt, actualInsertAt)
	}
}

func TestInsertAt_EmptyList(t *testing.T) {
	list := []int{}
	insertPoint := 2
	value := 7

	expectedInsertAt := []int{}
	expectedErr := definitions.ErrEmptyList

	actualInsertAt, actualErr := InsertAt(value, list, insertPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedInsertAt, actualInsertAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedInsertAt, actualInsertAt)
	}
}

func TestInsertAt_NegativeInsertPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	insertPoint := -2
	value := 7

	expectedInsertAt := []int{}
	expectedErr := definitions.ErrIndexOutOfRange

	actualInsertAt, actualErr := InsertAt(value, list, insertPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedInsertAt, actualInsertAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedInsertAt, actualInsertAt)
	}
}

func TestInsertAt_OutOfRangeInsertPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	insertPoint := 12
	value := 7

	expectedInsertAt := []int{}
	expectedErr := definitions.ErrIndexOutOfRange

	actualInsertAt, actualErr := InsertAt(value, list, insertPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedInsertAt, actualInsertAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedInsertAt, actualInsertAt)
	}
}

func TestInsertAt_ZeroInsertPoint(t *testing.T) {
	list := []int{1, 2, 3, 4}
	insertPoint := 0
	value := 7

	expectedInsertAt := []int{}
	expectedErr := definitions.ErrIndexOutOfRange

	actualInsertAt, actualErr := InsertAt(value, list, insertPoint)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedInsertAt, actualInsertAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedInsertAt, actualInsertAt)
	}
}

func TestRange(t *testing.T) {
	startValue, endValue := 4, 9

	expectedRange := []int{4, 5, 6, 7, 8, 9}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRange_SameStartValueAndEndValue(t *testing.T) {
	startValue, endValue := 4, 4

	expectedRange := []int{}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRange_NegativeStartValuePositiveEndValue(t *testing.T) {
	startValue, endValue := -4, 4

	expectedRange := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRange_NegativeStartValueNegativeEndValue(t *testing.T) {
	startValue, endValue := -7, -4

	expectedRange := []int{-7, -6, -5, -4}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRange_PositiveStartValueNegativeEndValue(t *testing.T) {
	startValue, endValue := 7, -4

	expectedRange := []int{7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRange_PositiveStartValueGreaterThanPositiveEndValue(t *testing.T) {
	startValue, endValue := 7, 4

	expectedRange := []int{7, 6, 5, 4}

	actualRange := Range(startValue, endValue)

	if !reflect.DeepEqual(expectedRange, actualRange) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRange, actualRange)
	}
}

func TestRndSelect(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	count := 3

	expectedRndSelectCount := 3
	var expectedErr error

	actualRndSelect, actualErr := RndSelect(list, count)
	actualRndSelectCount := len(actualRndSelect)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRndSelectCount, actualRndSelectCount) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRndSelectCount, actualRndSelectCount)
	}
}

func TestRndSelect_EmptyList(t *testing.T) {
	list := []int{}
	count := 3

	expectedRndSelect := []int{}
	expectedErr := definitions.ErrEmptyList

	actualRndSelect, actualErr := RndSelect(list, count)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRndSelect, actualRndSelect) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRndSelect, actualRndSelect)
	}
}

func TestRndSelect_NegativeCount(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	count := -3

	expectedRndSelect := []int{}
	expectedErr := definitions.ErrNegativeNumber

	actualRndSelect, actualErr := RndSelect(list, count)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRndSelect, actualRndSelect) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRndSelect, actualRndSelect)
	}
}

func TestRndSelect_ZeroCount(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	count := 0

	expectedRndSelect := []int{}
	var expectedErr error

	actualRndSelect, actualErr := RndSelect(list, count)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedRndSelect, actualRndSelect) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedRndSelect, actualRndSelect)
	}
}
