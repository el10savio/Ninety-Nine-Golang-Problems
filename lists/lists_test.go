package lists

import (
	"reflect"
	"testing"
)

func TestMyLast(t *testing.T) {
	list := []int{1, 2, 3, 4}

	expectedMyLast := 4
	var expectedErr error

	actualMyLast, actualErr := MyLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyLast, actualMyLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLast, actualMyLast)
	}
}

func TestMyLast_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyLast := 0
	expectedErr := ErrEmptyList

	actualMyLast, actualErr := MyLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyLast, actualMyLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLast, actualMyLast)
	}
}

func TestMyButLast(t *testing.T) {
	list := []int{1, 2, 3, 4}

	expectedMyButLast := 3
	var expectedErr error

	actualMyButLast, actualErr := MyButLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyButLast, actualMyButLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyButLast, actualMyButLast)
	}
}

func TestMyButLast_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyButLast := 0
	expectedErr := ErrAtLeastTwo

	actualMyButLast, actualErr := MyButLast(list)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedMyButLast, actualMyButLast) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyButLast, actualMyButLast)
	}
}

func TestElementAt(t *testing.T) {
	list := []int{1, 2, 3}
	index := 2

	expectedElementAt := 3
	var expectedErr error

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestElementAt_EmptyList(t *testing.T) {
	list := []int{}
	index := 2

	expectedElementAt := 0
	expectedErr := ErrEmptyList

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestElementAt_NegativeIndex(t *testing.T) {
	list := []int{1, 2, 3}
	index := -2

	expectedElementAt := 0
	expectedErr := ErrNegativeIndex

	actualElementAt, actualErr := ElementAt(list, index)

	if !reflect.DeepEqual(expectedErr, actualErr) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedErr, actualErr)
	}

	if !reflect.DeepEqual(expectedElementAt, actualElementAt) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedElementAt, actualElementAt)
	}
}

func TestMyLength(t *testing.T) {
	list := []int{1, 2, 3}

	expectedMyLength := 3

	actualMyLength := MyLength(list)

	if !reflect.DeepEqual(expectedMyLength, actualMyLength) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLength, actualMyLength)
	}
}

func TestMyLength_EmptyList(t *testing.T) {
	list := []int{}

	expectedMyLength := 0

	actualMyLength := MyLength(list)

	if !reflect.DeepEqual(expectedMyLength, actualMyLength) {
		t.Fatalf("Expected: %v\n Got: %v\n", expectedMyLength, actualMyLength)
	}
}
