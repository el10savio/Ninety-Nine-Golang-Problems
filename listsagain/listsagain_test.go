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
